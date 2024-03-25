package openpac

import (
	"encoding/json"

	"github.com/mitchellh/mapstructure"

	"github.com/zmicro-team/pab/utils"
)

func struct2Map(in any) (map[string]string, error) {
	b, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}
	m := make(map[string]string)
	err = json.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func concatMap(mp map[string]string) string {
	return utils.ConcatSortMap(mp, "=", "&") + "&"
}

func MapStructEncode(in any) (map[string]any, error) {
	b, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}
	m := make(map[string]any)
	err = json.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
	// TODO: bug 造成查询支付订单验名失败, 原因在于生成签名只加签字符串, 而以下生成类型
	// return mapstruct.Struct2MapTag(in, "json"), nil
}

func MapStructWeakDecode(input, output any) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook:       mapstructure.TextUnmarshallerHookFunc(),
		Result:           output,
		WeaklyTypedInput: true,
		TagName:          "json",
	})
	if err != nil {
		return err
	}
	return decoder.Decode(input)
}
