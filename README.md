# banking_spa
Simple single-page application written on Golang+MongoDB+React

## backend API
POST `/api/payments/from_card`
```
{
    card_number: "2424242424242424" | str, 16 digits
    card_expires: "MM/YY"           | str, YY between 19 and 35
    card_cvc: 776                   | int, between 100 and 999
    value: 7777                     | int, between 1000 and 75000
    comment: "txt"                  | str, length not more 150
    email: "example@mail.com"       | str, with email pattern
}
```

POST `/api/payments/via_bank`
```
{
    inn: "1234567890"                      | str, 10 or 12 digits
    bik: "123456789"                       | str, 9 digits
    account_number: "12345678901234567890" | str, 20 digits
    for_what: "без НДС"                    | str, "без НДС" or "НДС 10%" or "НДС 18%"
    value: 7777                            | int, between 1000 and 75000
}
```

POST `/api/payments/requests`
```
{
    inn: 1234567890                        | int, 10 or 12 digits
    bik: 123456789                         | int, 9 digits
    account_number: "12345678901234567890" | str, 20 digits
    for_what: "без НДС"                    | str, "без НДС" or "НДС 10%" or "НДС 18%"
    value: 7777                            | int, between 1000 and 75000
    phone: "+79997772211"                  | str, 10 digits
    email: "example@mail.com"              | str, with email pattern
}
```

PATCH `/api/payments/from_card/:paymentId`
```
{
    safe: false      | bool
}
```

GET `/api/payments/from_card?field=<field_name>&value=<value>`

GET `/api/payments/requests?field=<field_name>&value=<value>`

GET `/api/payments/from_card?field=<field_name>&desc=<bool_value>`

GET `/api/payments/requests?field=<field_name>&desc=<bool_value>`

GET `/api/copmanies/:companyId`
