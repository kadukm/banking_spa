# banking_spa
Simple single-page application written on Golang+MongoDB+React

#### How to run
```
npm run build
go run main.go
```

## TODO
- [ ] Send request to payment via bank using GET-form, not fetch
- [ ] Add vendoring
- [ ] Add opportunity to search in admin-panel by non-string fields
- [ ] Move font-files to another directory
- [ ] Add swagger
- [ ] Return payments in admin-panel by different pages
- [ ] Add description for all exported objects

## Database schema
You can find info about database in files:
- `./server/db/init.go` - database' and collections' names
- `./server/db/structs.go` - collections' description

## Backend API
POST `/api/payments/from_card`
```
{
    card_number: "2424242424242424" | str, 16 digits
    card_expires: "MM/YY"           | str, YY between 19 and 35
    card_cvc: "776"                 | str, 3 digits
    amount: 7777                    | int, between 1000 and 75000
    comment: "txt"                  | str, length not more 150
    email: "example@mail.com"       | str, with email pattern
}
```

POST `/api/payments/requests`
```
{
    inn: 1234567890                        | str, 10 or 12 digits
    bik: 123456789                         | str, 9 digits
    account_number: "12345678901234567890" | str, 20 digits
    for_what: "без НДС"                    | str, contains "без НДС" or "НДС 10%" or "НДС 18%"
    amount: 7777                           | int, between 1000 and 75000
    phone: "+79997772211"                  | str, 10 digits
    email: "example@mail.com"              | str, with email pattern
}
```

GET `/api/payments/via_bank?inn&bik&account_number&for_what&amount`

Fields in the query are the same as in `/api/payments/requests`

PATCH `/api/payments/from_card/:paymentID`
```
{
    dangerous: true      | bool
}
```

GET `/api/payments/from_card?field=<field_name>&value=<value>`

GET `/api/payments/requests?field=<field_name>&value=<value>`

GET `/api/payments/from_card/sort?field=<field_name>&desc=<bool_value>`

GET `/api/payments/requests/sort?field=<field_name>&desc=<bool_value>`

GET `/api/copmanies/:companyID`

GET `/api/companies/:companyID/products?maxcount=<int_value>`
