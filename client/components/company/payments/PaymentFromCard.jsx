import React from 'react'
import Modal from 'react-modal'
import '../../../styles/PaymentfromCard.css'
import apiBaseUrl from '../../../config.js'
import * as utils from '../../../utils.js'

export default class PaymentFromCard extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            modal: {
                show: false,
                message: undefined
            },
            ok: {
                card_number: undefined,
                card_expires: undefined,
                card_cvc: undefined,
                amount: undefined,
                comment: undefined,
                email: undefined
            },
            data: {
                card_number: '',
                card_expires: '',
                card_cvc: '',
                amount: '',
                amount: '',
                comment: '',
                email: ''
            }
        }
    }

    postPayment = () => {
        if (!this.isAllFieldsOk()) {
            this.setState({modal: {show: true, message: 'Одно или несколько полей заполнены некорректно'}})
            return
        }

        const init = {
            method: "POST",
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(this.state.data),
            mode: 'cors'
        }
        fetch(`${apiBaseUrl}/api/payments/from_card`, init)
            .then(response => response.json())
            .then(res => this.setState({modal: {show: true, message: res.result}}))
    }

    isAllFieldsOk = () => {
        return this.state.ok.card_number &&
               this.state.ok.card_expires &&
               this.state.ok.card_cvc &&
               this.state.ok.amount &&
               this.state.ok.email
    }

    closeModal = () => {
        this.setState({modal: {show: false, message: undefined}})
    }
    
    render() {
        return (
            <div>
                <Modal isOpen={this.state.modal.show}>
                    <div>
                        {this.state.modal.message}
                    </div>
                    <button onClick={this.closeModal}>Закрыть</button>
                </Modal>
                <div className="pay-card">
                    <div className="pay-card__card">
                        <div className="pay-card__card-bg">
                            <div className="pay-card__logos">
                                <img className="pay-card__logo"
                                    src="/assets/images/financial_services/visa.png"
                                    alt="visa logo"
                                />
                                <img className="pay-card__logo"
                                    src="/assets/images/financial_services/mastercard.png"
                                    alt="mastercard logo"
                                />
                                <img className="pay-card__logo"
                                    src="/assets/images/financial_services/maestro.png"
                                    alt="maestro logo"
                                />
                            </div>
                            <input className={`pay-card__card-number-input ${this.state.ok.card_number === false ? 'wrong-input' : ''}`}
                                type="text"
                                placeholder="Номер карты"
                                id="card_number"
                                onChange={this.onChangeCardNumber}
                                onBlur={this.onBlurCardNumber}
                                value={this.state.data.card_number}
                            />
                            <input className={`pay-card__small-input ${this.state.ok.card_expires === false ? 'wrong-input' : ''}`}
                                type="text"
                                placeholder="ММ/ГГ"
                                id="card_expires"
                                onChange={this.onChangeCardExpires}
                                onBlur={this.onBlurCardExpires}
                                value={this.state.data.card_expires}
                            /> 
                            <input className={`pay-card__small-input ${this.state.ok.card_cvc === false ? 'wrong-input' : ''}`}
                                type="text"
                                placeholder="CVC"
                                id="card_cvc"
                                onChange={this.onChangeCardCvc}
                                onBlur={this.onBlurCardCvc}
                                value={this.state.data.card_cvc}
                            />
                        </div>
                    </div>
                    <div className="pay-card__other-info">
                        <div>
                            <label htmlFor="amount">Сумма</label>
                            <input className={this.state.ok.amount === false ? 'wrong-input' : ''}
                                type="text"
                                placeholder="от 1 000 до 75 000₽"
                                name="amount"
                                id="amount"
                                onChange={this.onChangeAmount}
                                onBlur={this.onBlurAmount}
                                value={this.state.data.amount}
                            />
                        </div>
                        <div>
                            <label htmlFor="comment">Комментарий</label>
                            <input
                                type="text"
                                placeholder="До 150 символов"
                                name="comment"
                                id="comment"
                                maxLength="150"
                                onChange={this.onChangeComment}
                                value={this.state.data.comment}
                            />
                        </div>
                        <div>
                            <label htmlFor="email">Ваш email</label>
                            <input className={this.state.ok.email === false ? 'wrong-input' : ''}
                                type="email"
                                placeholder="Для квитанций об оплате"
                                name="email"
                                id="email"
                                onChange={this.onChangeEmail}
                                onBlur={this.onBlurEmail}
                                value={this.state.data.email}
                            />
                        </div>
                        <button onClick={this.postPayment}>Заплатить</button>
                    </div>
                </div>
            </div>
        )
    }

    onChangeCardNumber = (event) => {
        const value = utils.prepareCardNumber(event.target.value)
        this.setState({data: {...this.state.data, card_number: value}})
    }

    onChangeCardExpires = (event) => {
        const value = utils.prepareCardExpires(event.target.value)
        this.setState({data: {...this.state.data, card_expires: value}})
    }

    onChangeCardCvc = (event) => {
        const value = utils.prepareCardCvc(event.target.value)
        this.setState({data: {...this.state.data, card_cvc: value}})
    }

    onChangeAmount = (event) => {
        const value = utils.prepareAmount(event.target.value)
        this.setState({data: {...this.state.data, amount: value}})
    }

    onChangeComment = (event) => {
        const value = event.target.value
        this.setState({data: {...this.state.data, comment: value}})
    }

    onChangeEmail = (event) => {
        this.setState({data: {...this.state.data, email: event.target.value}})
    }

    onBlurCardNumber = () => {
        const cardNumberOk = utils.isCardNumberOk(this.state.data.card_number)
        this.setState({ok: {...this.state.ok, card_number: cardNumberOk}})
    }

    onBlurCardExpires = () => {
        const cardExpiresOk = utils.isCardExpiresOk(this.state.data.card_expires)
        this.setState({ok: {...this.state.ok, card_expires: cardExpiresOk}})
    }

    onBlurCardCvc = () => {
        const cardCvcOk = utils.isCardCvcOk(this.state.data.card_cvc)
        this.setState({ok: {...this.state.ok, card_cvc: cardCvcOk}})
    }

    onBlurAmount = () => {
        const amountOk = utils.isAmountOk(this.state.data.amount)
        this.setState({ok: {...this.state.ok, amount: amountOk}})
    }

    onBlurEmail = () => {
        const emailOk = utils.isEmailOk(this.state.data.email)
        this.setState({ok: {...this.state.ok, email: emailOk}})
    }

}