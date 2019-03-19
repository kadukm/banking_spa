import React from 'react';
import Requests from './Requests.jsx'
import PaymentsFromCard from './PaymentsFromCard.jsx'

export default class AdminPanel extends React.Component {
    constructor(props) {
        super(props)
        this.state = {isRequests: false}
    }
    render() {
        return (
            <div>
                <button onClick={() => this.setState({isRequests: false})}>
                    Платежи с карты
                </button>
                <button onClick={() => this.setState({isRequests: true})}>
                    Запрошенные платежи
                </button>
                {this.state.isRequests
                    ? <Requests />
                    : <PaymentsFromCard />}
            </div>
        )
    }
}
