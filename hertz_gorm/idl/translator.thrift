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
    1: string translation
    2: string ukphone
    3: string usphone
    4: list<string> translations
    5: list<string> word_forms
    6: list<Etymology> etymologies
    7: list<EGSentence> eg_sentences
}

service TranslatorService {
    TranslateResponse TranslateText(1: TranslateRequest req) (api.get="/v1/translate")
}
