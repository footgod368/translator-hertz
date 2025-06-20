package service

import (
	"context"
	"github.com/bytedance/gg/gslice"
	"github.com/cloudwego/hertz-examples/bizdemo/hertz_gorm/biz/service/youdao"

	"github.com/cloudwego/hertz-examples/bizdemo/hertz_gorm/biz/hertz_gen/translator"
)

func Translate(ctx context.Context, req translator.TranslateRequest) (*translator.TranslateResponse, error) {
	youDaoResp, err := youdao.Query(ctx, req.Text)
	if err != nil {
		return nil, err
	}
	return convYouDaoResp(youDaoResp), nil
}

func convYouDaoResp(youDaoResp *youdao.YouDaoResponse) *translator.TranslateResponse {
	if youDaoResp == nil {
		return nil
	}
	wordData := youDaoResp.Ec.Word[0]

	return &translator.TranslateResponse{
		Ukphone:      wordData.Ukphone,
		Usphone:      wordData.Usphone,
		Translations: gslice.Map(wordData.Trs, func(tr youdao.Translation) string { return tr.Tr[0].L.I[0] }),
		WordForms:    gslice.Map(wordData.Wfs, func(wf youdao.WordFrom) string { return wf.Wf.Name + wf.Wf.Value }),
		Etymologies: gslice.Map(youDaoResp.Etym.Etyms.Zh, func(etymology youdao.EtymologyZh) *translator.Etymology {
			return &translator.Etymology{Value: etymology.Value, Desc: etymology.Desc}
		}),
		EgSentences: gslice.Map(youDaoResp.BlngSentsPart.SentencePair, func(sentencePair youdao.EgSentencePair) *translator.EGSentence {
			return &translator.EGSentence{Sentence: sentencePair.Sentence, Translation: sentencePair.SentenceTranslation}
		}),
	}
}
