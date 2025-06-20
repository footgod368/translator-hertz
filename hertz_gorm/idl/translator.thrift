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

service TranslatorService {
    TranslateResponse TranslateText(1: TranslateRequest req) (api.get="/v1/translate")
}
