// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package lang

import (
	"github.com/hashicorp/hcl/v2/ext/tryfunc"
	ctyyaml "github.com/zclconf/go-cty-yaml"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
	"github.com/zclconf/go-cty/cty/function/stdlib"

	"github.com/terraform-linters/tflint/terraform/collections"
	"github.com/terraform-linters/tflint/terraform/lang/funcs"
	"github.com/terraform-linters/tflint/terraform/lang/funcs/terraform"
)

var impureFunctions = []string{
	"bcrypt",
	"timestamp",
	"uuid",
}

// filesystemFunctions are the functions that allow interacting with arbitrary
// paths in the local filesystem, and which can therefore have their results
// vary based on something other than their arguments, and might allow template
// rendering to expose details about the system where Terraform is running.
var filesystemFunctions = collections.NewSetCmp[string](
	"file",
	"fileexists",
	"fileset",
	"filebase64",
	"filebase64sha256",
	"filebase64sha512",
	"filemd5",
	"filesha1",
	"filesha256",
	"filesha512",
	"templatefile",
)

// templateFunctions are functions that render nested templates. These are
// callable from module code but not from within the templates they are
// rendering.
var templateFunctions = collections.NewSetCmp[string](
	"templatefile",
	"templatestring",
)

// Functions returns the set of functions that should be used to when evaluating
// expressions in the receiving scope.
func (s *Scope) Functions() map[string]function.Function {
	s.funcsLock.Lock()
	if s.funcs == nil {
		// Some of our functions are just directly the cty stdlib functions.
		// Others are implemented in the subdirectory "funcs" here in this
		// repository. New functions should generally start out their lives
		// in the "funcs" directory and potentially graduate to cty stdlib
		// later if the functionality seems to be something domain-agnostic
		// that would be useful to all applications using cty functions.
		//
		// If you're adding something here, please consider whether it meets
		// the criteria for either or both of the sets [filesystemFunctions]
		// and [templateFunctions] and add it there if so, to ensure that
		// functions relying on those classifications will behave correctly.
		coreFuncs := map[string]function.Function{
			"abs":              stdlib.AbsoluteFunc,
			"abspath":          funcs.AbsPathFunc,
			"alltrue":          funcs.AllTrueFunc,
			"anytrue":          funcs.AnyTrueFunc,
			"basename":         funcs.BasenameFunc,
			"base64decode":     funcs.Base64DecodeFunc,
			"base64encode":     funcs.Base64EncodeFunc,
			"base64gzip":       funcs.Base64GzipFunc,
			"base64sha256":     funcs.Base64Sha256Func,
			"base64sha512":     funcs.Base64Sha512Func,
			"bcrypt":           funcs.BcryptFunc,
			"can":              tryfunc.CanFunc,
			"ceil":             stdlib.CeilFunc,
			"chomp":            stdlib.ChompFunc,
			"cidrhost":         funcs.CidrHostFunc,
			"cidrnetmask":      funcs.CidrNetmaskFunc,
			"cidrsubnet":       funcs.CidrSubnetFunc,
			"cidrsubnets":      funcs.CidrSubnetsFunc,
			"coalesce":         funcs.CoalesceFunc,
			"coalescelist":     stdlib.CoalesceListFunc,
			"compact":          stdlib.CompactFunc,
			"concat":           stdlib.ConcatFunc,
			"contains":         stdlib.ContainsFunc,
			"csvdecode":        stdlib.CSVDecodeFunc,
			"dirname":          funcs.DirnameFunc,
			"distinct":         stdlib.DistinctFunc,
			"element":          stdlib.ElementFunc,
			"endswith":         funcs.EndsWithFunc,
			"ephemeralasnull":  funcs.EphemeralAsNullFunc,
			"chunklist":        stdlib.ChunklistFunc,
			"file":             funcs.MakeFileFunc(s.BaseDir, false),
			"fileexists":       funcs.MakeFileExistsFunc(s.BaseDir),
			"fileset":          funcs.MakeFileSetFunc(s.BaseDir),
			"filebase64":       funcs.MakeFileFunc(s.BaseDir, true),
			"filebase64sha256": funcs.MakeFileBase64Sha256Func(s.BaseDir),
			"filebase64sha512": funcs.MakeFileBase64Sha512Func(s.BaseDir),
			"filemd5":          funcs.MakeFileMd5Func(s.BaseDir),
			"filesha1":         funcs.MakeFileSha1Func(s.BaseDir),
			"filesha256":       funcs.MakeFileSha256Func(s.BaseDir),
			"filesha512":       funcs.MakeFileSha512Func(s.BaseDir),
			"flatten":          stdlib.FlattenFunc,
			"floor":            stdlib.FloorFunc,
			"format":           stdlib.FormatFunc,
			"formatdate":       stdlib.FormatDateFunc,
			"formatlist":       stdlib.FormatListFunc,
			"indent":           stdlib.IndentFunc,
			"index":            funcs.IndexFunc, // stdlib.IndexFunc is not compatible
			"join":             stdlib.JoinFunc,
			"jsondecode":       stdlib.JSONDecodeFunc,
			"jsonencode":       stdlib.JSONEncodeFunc,
			"keys":             stdlib.KeysFunc,
			"length":           funcs.LengthFunc,
			"list":             funcs.ListFunc,
			"log":              stdlib.LogFunc,
			"lookup":           funcs.LookupFunc,
			"lower":            stdlib.LowerFunc,
			"map":              funcs.MapFunc,
			"matchkeys":        funcs.MatchkeysFunc,
			"max":              stdlib.MaxFunc,
			"md5":              funcs.Md5Func,
			"merge":            stdlib.MergeFunc,
			"min":              stdlib.MinFunc,
			"one":              funcs.OneFunc,
			"parseint":         stdlib.ParseIntFunc,
			"pathexpand":       funcs.PathExpandFunc,
			"plantimestamp":    funcs.PlantimestampFunc,
			"pow":              stdlib.PowFunc,
			"range":            stdlib.RangeFunc,
			"regex":            stdlib.RegexFunc,
			"regexall":         stdlib.RegexAllFunc,
			"replace":          funcs.ReplaceFunc,
			"reverse":          stdlib.ReverseListFunc,
			"rsadecrypt":       funcs.RsaDecryptFunc,
			"sensitive":        funcs.SensitiveFunc,
			"nonsensitive":     funcs.NonsensitiveFunc,
			"issensitive":      funcs.IssensitiveFunc,
			"setintersection":  stdlib.SetIntersectionFunc,
			"setproduct":       stdlib.SetProductFunc,
			"setsubtract":      stdlib.SetSubtractFunc,
			"setunion":         stdlib.SetUnionFunc,
			"sha1":             funcs.Sha1Func,
			"sha256":           funcs.Sha256Func,
			"sha512":           funcs.Sha512Func,
			"signum":           stdlib.SignumFunc,
			"slice":            stdlib.SliceFunc,
			"sort":             stdlib.SortFunc,
			"split":            stdlib.SplitFunc,
			"startswith":       funcs.StartsWithFunc,
			"strcontains":      funcs.StrContainsFunc,
			"strrev":           stdlib.ReverseFunc,
			"substr":           stdlib.SubstrFunc,
			"sum":              funcs.SumFunc,
			"textdecodebase64": funcs.TextDecodeBase64Func,
			"textencodebase64": funcs.TextEncodeBase64Func,
			"timestamp":        funcs.TimestampFunc,
			"timeadd":          stdlib.TimeAddFunc,
			"timecmp":          funcs.TimeCmpFunc,
			"title":            stdlib.TitleFunc,
			"tostring":         funcs.MakeToFunc(cty.String),
			"tonumber":         funcs.MakeToFunc(cty.Number),
			"tobool":           funcs.MakeToFunc(cty.Bool),
			"toset":            funcs.MakeToFunc(cty.Set(cty.DynamicPseudoType)),
			"tolist":           funcs.MakeToFunc(cty.List(cty.DynamicPseudoType)),
			"tomap":            funcs.MakeToFunc(cty.Map(cty.DynamicPseudoType)),
			"transpose":        funcs.TransposeFunc,
			"trim":             stdlib.TrimFunc,
			"trimprefix":       stdlib.TrimPrefixFunc,
			"trimspace":        stdlib.TrimSpaceFunc,
			"trimsuffix":       stdlib.TrimSuffixFunc,
			"try":              tryfunc.TryFunc,
			"upper":            stdlib.UpperFunc,
			"urlencode":        funcs.URLEncodeFunc,
			"uuid":             funcs.UUIDFunc,
			"uuidv5":           funcs.UUIDV5Func,
			"values":           stdlib.ValuesFunc,
			"yamldecode":       ctyyaml.YAMLDecodeFunc,
			"yamlencode":       ctyyaml.YAMLEncodeFunc,
			"zipmap":           stdlib.ZipmapFunc,
		}

		// Our two template-rendering functions want to be able to call
		// all of the other functions themselves, but we pass them indirectly
		// via a callback to avoid chicken/egg problems while initializing
		// the functions table.
		funcsFunc := func() (funcs map[string]function.Function, fsFuncs collections.Set[string], templateFuncs collections.Set[string]) {
			// The templatefile and templatestring functions prevent recursive
			// calls to themselves and each other by copying this map and
			// overwriting the relevant entries.
			return s.funcs, filesystemFunctions, templateFunctions
		}
		coreFuncs["templatefile"] = funcs.MakeTemplateFileFunc(s.BaseDir, funcsFunc)
		coreFuncs["templatestring"] = funcs.MakeTemplateStringFunc(funcsFunc)

		if s.PureOnly {
			// Force our few impure functions to return unknown so that we
			// can defer evaluating them until a later pass.
			for _, name := range impureFunctions {
				coreFuncs[name] = function.Unpredictable(s.funcs[name])
			}
		}

		// All of the built-in functions are also available under the "core::"
		// namespace, to distinguish from the "provider::" and "module::"
		// namespaces that can serve as external extension points.
		s.funcs = make(map[string]function.Function, len(coreFuncs)*2)
		for name, fn := range coreFuncs {
			s.funcs[name] = fn
			s.funcs["core::"+name] = fn
		}

		// Built-in Terraform provider-defined functions are typically obtained dynamically,
		// but given that they are built-ins, they are provided just like regular functions.
		s.funcs["provider::terraform::encode_tfvars"] = terraform.EncodeTfvarsFunc
		s.funcs["provider::terraform::decode_tfvars"] = terraform.DecodeTfvarsFunc
		s.funcs["provider::terraform::encode_expr"] = terraform.EncodeExprFunc
	}
	s.funcsLock.Unlock()

	return s.funcs
}

// NewMockFunction creates a mock function that returns a dynamic value.
// This is primarily used to replace provider-defined functions.
func NewMockFunction(call *FunctionCall) function.Function {
	return function.New(&function.Spec{
		VarParam: &function.Parameter{
			Type:             cty.DynamicPseudoType,
			AllowNull:        true,
			AllowUnknown:     true,
			AllowDynamicType: true,
			AllowMarked:      true,
		},
		Type: function.StaticReturnType(cty.DynamicPseudoType),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			return cty.DynamicVal, nil
		},
	})
}
