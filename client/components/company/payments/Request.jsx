import React from 'react';
import '../../../styles/Request.css'
import apiBaseUrl from '../../../config.js'
import * as utils from '../../../utils.js'

export default class Request extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            ok: {
                inn: undefined,
                bik: undefined,
                account_number: undefined,
                for_what: undefined,
                amount: undefined,
                phone: undefined,
                email: undefined
            },
            data: {
                inn: '',
                bik: '',
                account_number: '',
                for_what: '',
                amount: '',
                phone: '',
                email: ''
            }
        }
    }

    postRequest = () => {
        //TODO: check values
        let init = {
            method: "POST",
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(this.state.data),
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
                        Создайте платежку, а {this.props.status} {this.props.name} подпишет её у себя в интернет-банке
                    </strong>
                </header>
                <div>
                    <div>
                        <label htmlFor="inn">ИНН получателя</label>
                        <input className={this.state.ok.inn === false ? 'wrong-input' : ''}
                            type="text"
                            name="inn"
                            id="inn"
                            placeholder="ИНН (10 цифр)"
                            onChange={this.onChangeInn}
                            onBlur={this.onBlurInn}
                            value={this.state.data.inn}
                        />
                    </div>
                    <div>
                        <label htmlFor="bik">БИК</label>
                        <input className={this.state.ok.bik === false ? 'wrong-input' : ''}
                            type="text"
                            name="bik"
                            id="bik"
                            placeholder="БИК (9 цифр)"
                            onChange={this.onChangeBik}
                            onBlur={this.onBlurBik}
                            value={this.state.data.bik}
                        />
                    </div>
                    <div>
                        <label htmlFor="account_number">Номер счёта</label>
                        <input className={this.state.ok.account_number === false ? 'wrong-input' : ''}
                            type="text"
                            name="account_number"
                            id="account_number"
                            placeholder="Номер счёта (20 цифр)"
                            onChange={this.onChangeAccountNumber}
                            onBlur={this.onBlurAccountNumber}
                            value={this.state.data.account_number}
                        />
                    </div>
                    <div>
                        <label htmlFor="for_what">За что</label>
                        <input className={this.state.ok.for_what === false ? 'wrong-input' : ''}
                            type="text"
                            id="for_what"
                            name="for_what"
                            onChange={this.onChangeForWhat}
                            onBlur={this.onBlurForWhat}
                            value={this.state.data.for_what}
                        />
                    </div>
                    <div>
                        <label htmlFor="amount">Сколько</label>
                        <input className={this.state.ok.amount === false ? 'wrong-input' : ''}
                            type="text"
                            name="amount"
                            id="amount"
                            placeholder="От 1000 до 75000 рублей"
                            onChange={this.onChangeAmount}
                            onBlur={this.onBlurAmount}
                            value={this.state.data.amount}
                        />
                    </div>
                    <div>
                        <label htmlFor="phone">Телефон</label>
                        <input className={this.state.ok.phone === false ? 'wrong-input' : ''}
                            type="tel"
                            name="phone"
                            id="phone"
                            placeholder="+79999999999"
                            onChange={this.onChangePhone}
                            onBlur={this.onBlurPhone}
                            value={this.state.data.phone}
                        />
                    </div>
                    <div>
                        <label htmlFor="email">Ваш email</label>
                        <input className={this.state.ok.email === false ? 'wrong-input' : ''}
                            type="email"
                            name="email"
                            id="email"
                            onChange={this.onChangeEmail}
                            onBlur={this.onBlurEmail}
                            value={this.state.data.email}
                       />
                    </div>
                    <button onClick={this.postRequest}>Заплатить</button>
                </div>
            </div>
        )
    }

    onChangeInn = (event) => {
        const value = utils.prepareInn(event.target.value)
        this.setState({data: {...this.state.data, inn: value}})
    }

    onChangeBik = (event) => {
        const value = utils.prepareBik(event.target.value)
        this.setState({data: {...this.state.data, bik: value}})
    }

    onChangeAccountNumber = (event) => {
        const value = utils.prepareAccountNumber(event.target.value)
        this.setState({data: {...this.state.data, account_number: value}})
    }

    onChangeForWhat = (event) => {
        this.setState({data: {...this.state.data, for_what: event.target.value}})
    }

    onChangeAmount = (event) => {
        const value = utils.prepareAmount(event.target.value)
        this.setState({data: {...this.state.data, amount: value}})
    }

    onChangePhone = (event) => {
        const value = utils.preparePhone(event.target.value)
        this.setState({data: {...this.state.data, phone: value}})
    }

    onChangeEmail = (event) => {
        this.setState({data: {...this.state.data, email: event.target.value}})
    }

    onBlurInn = () => {
        const innOk = utils.isInnOk(this.state.data.inn)
        this.setState({ok: {...this.state.ok, inn: innOk}})
    }

    onBlurBik = () => {
        const bikOk = utils.isBikOk(this.state.data.bik)
        this.setState({ok: {...this.state.ok, bik: bikOk}})
    }

    onBlurAccountNumber = () => {
        const accountNumberOk = utils.isAccountNumberOk(this.state.data.account_number)
        this.setState({ok: {...this.state.ok, account_number: accountNumberOk}})
    }

    onBlurForWhat = () => {
        const forWhatOk = utils.isForWhatOk(this.state.data.for_what)
        this.setState({ok: {...this.state.ok, for_what: forWhatOk}})
    }

    onBlurAmount = () => {
        const amountOk = utils.isAmountOk(this.state.data.amount)
        this.setState({ok: {...this.state.ok, amount: amountOk}})
    }

    onBlurPhone = () => {
        const phoneOk = utils.isPhoneOk(this.state.data.phone)
        this.setState({ok: {...this.state.ok, phone: phoneOk}})
    }

    onBlurEmail = () => {
        const emailOk = utils.isEmailOk(this.state.data.email)
        this.setState({ok: {...this.state.ok, email: emailOk}})
    }
}