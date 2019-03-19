import React from 'react';
import Requests from './Requests.jsx'
import PaymentsFromCard from './PaymentsFromCard.jsx'
import '../../styles/AdminPanel.css'

export default class AdminPanel extends React.Component {
    constructor(props) {
        super(props)
        this.state = {isRequests: false}
    }
    render() {
        return (
            <div>
                <button
                    className=""
                    onClick={() => this.setState({isRequests: false})}
                >
                    Платежи с карты
                </button>
                <button
                    className=""
                    onClick={() => this.setState({isRequests: true})}
                >
                    Запрошенные платежи
                </button>
                {this.state.isRequests
                    ? <Requests />
                    : <PaymentsFromCard />}
            </div>
        )
    }
}
