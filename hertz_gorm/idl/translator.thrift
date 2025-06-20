namespace go translator

struct TranslateRequest {
    1: string text
}

struct TranslateResponse {
    1: string translation
}

service TranslatorService {
    TranslateResponse TranslateText(1: TranslateRequest req) (api.get="/v1/translate")
}
