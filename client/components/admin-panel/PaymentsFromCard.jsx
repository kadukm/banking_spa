import React from 'react';
import Modal from 'react-modal'
import apiBaseUrl from '../../config.js'
import {TableKey, TableCell} from '../utils/TableItems.jsx'

const paymentFromCardApiPath = `${apiBaseUrl}/api/payments/from_card`
const paymentFields = ['id', 'card_number', 'card_expires', 'card_cvc', 'amount', 'comment', 'email', 'dangerous']

export default class PaymentsFromCard extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            modal: {
                show: false,
                message: undefined
            },
            payments: null,
            sort: {
                field: '',
                desc: false
            },
            filter: {
                field: '',
                value: ''
            }
        }
    }

    componentDidMount() {
        this.updatePaymentsTable(paymentFromCardApiPath)
    }

    closeModal = () => {
        this.setState({modal: {show: false, message: undefined}})
    }

    buildUrlForSort = () => {
        const query = this.state.sort.field == '' ? '' : `?field=${this.state.sort.field}&desc=${this.state.sort.desc}`
        return `${paymentFromCardApiPath}/sort${query}`
    }

    buildUrlForFilter = () => {
        const query = this.state.filter.field == '' ? '' : `?field=${this.state.filter.field}&value=${encodeURIComponent(this.state.filter.value)}`
        return paymentFromCardApiPath + query
    }

    getSortedPayments = () => {
        const url = this.buildUrlForSort()
        this.updatePaymentsTable(url)
    }

    getFilteredPayments = () => {
        const url = this.buildUrlForFilter()
        this.updatePaymentsTable(url)
    }

    updatePaymentsTable = (url) => {
        fetch(url)
            .then(response => response.json())
            .then(res => {
                if (res.ok) {
                    this.setState({payments: res.result})
                } else {
                    this.setState({modal: {show: true, message: res.result}})
                }
            })
    }

    patchPayment = (paymentIdx) => {
        const paymentID = this.state.payments[paymentIdx].id
        const paymentDangerousNewValue = !this.state.payments[paymentIdx].dangerous
        const init = {
            method: "PATCH",
            body: JSON.stringify({dangerous: paymentDangerousNewValue})
        }
        fetch(`${paymentFromCardApiPath}/${paymentID}`, init)
            .then(response => response.json())
            .then(res => {
                if (res.ok) {
                    this.state.payments[paymentIdx].dangerous = paymentDangerousNewValue
                    this.setState({payments: this.state.payments})
                } else {
                    this.setState({modal: {show: true, message: res.result}})
                }
            })
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
                <div>
                    <div>
                        <input
                            type="text"
                            placeholder="Поле"
                            value={this.state.sort.field}
                            onChange={(e) => this.setState({sort: {...this.state.sort, field: e.target.value}})}
                            id="sort_field"
                        />
                        <label htmlFor="sort_desc">По убыванию: </label>
                        <input
                            type="checkbox"
                            checked={this.state.sort.desc}
                            onChange={(e) => this.setState({sort: {...this.state.sort, desc: e.target.checked}})}
                            id="sort_desc"
                        />
                        <button onClick={this.getSortedPayments}>Сортировать</button>
                    </div>
                    <div>
                        <input
                            type="text"
                            placeholder="Поле"
                            value={this.state.filter.field}
                            onChange={(e) => this.setState({filter: {...this.state.filter, field: e.target.value}})}
                            id="filter_field"
                        />
                        <input
                            type="text"
                            placeholder="Значение"
                            value={this.state.filter.value}
                            onChange={(e) => this.setState({filter: {...this.state.filter, value: e.target.value}})}
                            id="filter_value"
                        />
                        <button onClick={this.getFilteredPayments}>Отфильтровать</button>
                    </div>
                </div>
                {this.state.payments !== null && (
                    <div className="table">
                        <div className="table__row">
                            {
                                paymentFields.map(fieldName => (
                                    <TableKey
                                        name={fieldName}
                                        key={fieldName}
                                    />
                                ))
                            }
                        </div>
                        {
                            this.state.payments.map((payment, idx) => (
                                <div className={"table__row" + (payment.dangerous ? " dangerous" : "")} key={payment.id}>
                                    {
                                        paymentFields
                                            .slice(0, -1)
                                            .map(fieldName => (
                                                <TableCell key={`${payment.id}__${fieldName}`}>
                                                    {payment[fieldName]}
                                                </TableCell>
                                            ))
                                    }
                                    <TableCell key={`${payment.id}__dangerous`}>
                                        <input
                                            type="checkbox"
                                            onChange={() => this.patchPayment(idx)}
                                            checked={this.state.payments[idx].dangerous}
                                        />
                                    </TableCell>
                                </div>
                            ))
                        }
                    </div>
                )}
            </div>
        )
    }
}
