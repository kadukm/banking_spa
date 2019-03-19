import React from 'react';
import apiBaseUrl from '../../config.js'
import {TableKey, TableCell} from '../utils/TableItems.jsx'

const paymentRequestApiPath = `${apiBaseUrl}/api/payments/requests`
const paymentFields = ['id', 'inn', 'bik', 'account_number', 'for_what', 'amount', 'phone', 'email']

export default class Requests extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
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
        this.updatePaymentsTable(paymentRequestApiPath)
    }

    buildUrlForSort = () => {
        const query = this.state.sort.field == '' ? '' : `?field=${this.state.sort.field}&desc=${this.state.sort.desc}`
        return `${paymentRequestApiPath}/sort${query}`
    }

    buildUrlForFilter = () => {
        const query = this.state.filter.field == '' ? '' : `?field=${this.state.filter.field}&value=${encodeURIComponent(this.state.filter.value)}`
        return paymentRequestApiPath + query
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
                    //TODO: show modal
                }
            })
    }

    render() {
        return (
            <div>
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
                                            .map(fieldName => (
                                                <TableCell key={`${payment.id}__${fieldName}`}>
                                                    {payment[fieldName]}
                                                </TableCell>
                                            ))
                                    }
                                </div>
                            ))
                        }
                    </div>
                )}
            </div>
        )
    }
}
