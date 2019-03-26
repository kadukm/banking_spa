import React from 'react';
import Modal from 'react-modal'
import '../../../styles/PaymentViaBank.css'
import apiBaseUrl from '../../../config.js'
import * as utils from '../../../utils.js'

export default class PaymentViaBank extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            modal: {
                show: false,
                message: undefined
            },
            ok: {
                inn: undefined,
                bik: undefined,
                account_number: undefined,
                for_what: undefined,
                amount: undefined
            },
            data: {
                inn: '',
                bik: '',
                account_number: '',
                for_what: '',
                amount: ''
            }
        }
    }

    getPaymentViaBank = () => {
        if (!this.isAllFieldsOk()) {
            this.setState({modal: {show: true, message: 'Одно или несколько полей не заполнены или заполнены некорректно'}})
            return
        }

        const query = utils.urlEncodeObject(this.state.data)
        const init = {
            method: "GET",
            mode: 'cors'
        }
        fetch(`${apiBaseUrl}/api/payments/via_bank?${query}`, init)
            .then(response => {
                if (response.status === 200) {
                    response.blob()
                        .then(blob => {
                            var a = document.createElement('a');
                            a.href = window.URL.createObjectURL(blob);
                            a.download = 'payment.pdf';
                            document.body.appendChild(a);
                            a.click();
                            document.body.removeChild(a);
                        })
                } else {
                    response.json()
                        .then(res => this.setState({modal: {show: true, message: res.result}}))   
                }
            })
    }

    isAllFieldsOk = () => {
        return this.state.ok.inn &&
               this.state.ok.bik &&
               this.state.ok.account_number &&
               this.state.ok.for_what &&
               this.state.ok.amount
    }

    closeModal = () => {
        this.setState({modal: {show: false, message: undefined}})
    }

    render() {
        return (
            <div className="payment-via-bank">
                <Modal isOpen={this.state.modal.show}>
                    <div>
                        {this.state.modal.message}
                    </div>
                    <button onClick={this.closeModal}>Закрыть</button>
                </Modal>
                <header>
                    <strong>
                        Сформируйте платёжку и загрузите её в свой банк для подписи
                    </strong>
                </header>
                <div>
                    <div className="standard-field">
                        <label htmlFor="inn">От кого</label>
                        <input className={`standard-input ${this.state.ok.inn === false ? 'wrong-input' : ''}`}
                            type="text"
                            name="inn"
                            id="inn"
                            placeholder="ИНН (10 цифр)"
                            onChange={this.onChangeInn}
                            onBlur={this.onBlurInn}
                            value={this.state.data.inn}
                        />
                    </div>
                    <div className="standard-field">
                        <label htmlFor="bik">БИК</label>
                        <input className={`standard-input ${this.state.ok.bik === false ? 'wrong-input' : ''}`}
                            type="text"
                            name="bik"
                            id="bik"
                            placeholder="БИК (9 цифр)"
                            onChange={this.onChangeBik}
                            onBlur={this.onBlurBik}
                            value={this.state.data.bik}
                        />
                    </div>
                    <div className="standard-field">
                        <label htmlFor="account_number">Номер счёта</label>
                        <input className={`standard-input ${this.state.ok.account_number === false ? 'wrong-input' : ''}`}
                            type="text"
                            name="account_number"
                            id="account_number"
                            placeholder="Номер счёта (20 цифр)"
                            onChange={this.onChangeAccountNumber}
                            onBlur={this.onBlurAccountNumber}
                            value={this.state.data.account_number}
                        />
                    </div>
                    <div className="standard-field">
                        <label htmlFor="for_what">За что</label>
                        <input className={`standard-input ${this.state.ok.for_what === false ? 'wrong-input' : ''}`}
                            type="text"
                            id="for_what"
                            name="for_what"
                            onChange={this.onChangeForWhat}
                            onBlur={this.onBlurForWhat}
                            value={this.state.data.for_what}
                        />
                        <div className="nds-options">
                            <button onClick={() => this.onClickNdsButton("НДС 18%")}>НДС 18%</button>
                            <button onClick={() => this.onClickNdsButton("НДС 10%")}>НДС 10%</button>
                            <button onClick={() => this.onClickNdsButton("без НДС")}>без НДС</button>
                        </div>
                    </div>
                    <div className="standard-field">
                        <label htmlFor="amount">Сколько</label>
                        <input className={`standard-input ${this.state.ok.amount === false ? 'wrong-input' : ''}`}
                            type="text"
                            name="amount"
                            id="amount"
                            placeholder="От 1000 до 75000 рублей"
                            onChange={this.onChangeAmount}
                            onBlur={this.onBlurAmount}
                            value={this.state.data.amount}
                        />
                    </div>
                    <button className="button-tochka"
                        onClick={this.getPaymentViaBank}
                    >
                        Получить файл для интернет-банка
                    </button>
                </div>
            </div>
        )
    }

    onClickNdsButton = (ndsValue) => {
        const value = utils.handleNds(ndsValue, this.state.data.for_what)
        this.setState({data: {...this.state.data, for_what: value}})
        const forWhatOk = utils.isForWhatOk(value)
        this.setState({ok: {...this.state.ok, for_what: forWhatOk}})
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
}