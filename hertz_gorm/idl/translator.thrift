namespace go translator

struct TranslateRequest {
    1: string text
}

struct Etymology {
    1: string value
    2: string desc
}

struct EGSentence {
    1: string sentence
    2: string translation
}

struct TranslateResponse {
    1: string ukphone
    2: string usphone
    3: list<string> translations
    4: list<string> word_forms
    5: list<Etymology> etymologies
    6: list<EGSentence> eg_sentences
}

struct CompleteRequest {
    1: string prefix
}

struct CompleteResponse {
    1: list<string> suggestions
}

struct DuckSayRequest {
}

struct DuckSayResponse {
    1: string text
}

struct SynonymsRequest {
    1: string word
}

struct SynonymsResponse {
    1: string synonyms_md
}

service TranslatorService {
    TranslateResponse TranslateText(1: TranslateRequest req) (api.get="/v1/translate")
    CompleteResponse Complete(1: CompleteRequest req) (api.get="/v1/complete")
    DuckSayResponse DuckSay(1: DuckSayRequest req) (api.get="/v1/duck_say")
    SynonymsResponse Synonyms(1: SynonymsRequest req) (api.get="/v1/synonyms")
}
