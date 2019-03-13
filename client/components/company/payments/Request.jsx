import React from 'react';
import apiBaseUrl from '../../../config.js'

export default class Request extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            inn: undefined,
            bik: undefined,
            account_number: undefined,
            for_what: undefined,
            amount: undefined,
            phone: undefined,
            email: undefined
        }
    }

    postRequest = () => {
        //TODO: check values
        let init = {
            method: "POST",
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(this.state),
            mode: 'cors'
        }
        fetch(`${apiBaseUrl}/api/payments/requests`, init)
            .then(response => response.json)
            .then(res => console.log(res))
            //TODO: show user that all's goods
    }

    render() {
        return (
            <div>
                <header>
                    <strong>
                        Создайте платежку, а {this.props.status} {this.props.name} подпишет её у себя в
                        интернет-банке
                    </strong>
                </header>
                <div>
                    <div>
                        <label htmlFor="inn">ИНН получателя</label>
                        <input type="text" name="inn" id="inn" onChange={this.onChangeInn} />
                    </div>
                    <div>
                        <label htmlFor="bik">БИК</label>
                        <input type="text" name="bik" id="bik" onChange={this.onChangeBik} />
                    </div>
                    <div>
                        <label htmlFor="account_number">Номер счёта</label>
                        <input type="text" name="account_number" id="account_number" onChange={this.onChangeAccountNumber} />
                    </div>
                    <div>
                        <label htmlFor="for_what">За что</label>
                        <input type="text" id="for_what" name="for_what" onChange={this.onChangeForWhat} />
                    </div>
                    <div>
                        <label htmlFor="amount">Сколько</label>
                        <input type="text" name="amount" id="amount" onChange={this.onChangeAmount} />
                    </div>
                    <div>
                        <label htmlFor="phone">Телефон</label>
                        <input type="tel" name="phone" id="phone" onChange={this.onChangePhone} />
                    </div>
                    <div>
                        <label htmlFor="email">Ваш email</label>
                        <input type="email" name="email" id="email" onChange={this.onChangeEmail} />
                    </div>
                    <button onClick={this.postRequest}>Заплатить</button>
                </div>
            </div>
        )
    }
    
    onChange = (event, fieldname) => {
        this.setState({[fieldname]: event.target.value})
    }

    onChangeInn = (event) => {
        this.onChange(event, 'inn')
    }

    onChangeBik = (event) => {
        this.onChange(event, 'bik')
    }

    onChangeAccountNumber = (event) => {
        this.onChange(event, 'account_number')
    }

    onChangeForWhat = (event) => {
        this.onChange(event, 'for_what')
    }

    onChangeAmount = (event) => {
        this.setState({amount: parseInt(event.target.value)})
    }

    onChangePhone = (event) => {
        this.onChange(event, 'phone')
    }

    onChangeEmail = (event) => {
        this.onChange(event, 'email')
    }
}