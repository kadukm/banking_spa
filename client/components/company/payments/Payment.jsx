import React from 'react';
import PaymentFromCard from './PaymentFromCard.jsx'
import PaymentViaBank from './PaymentViaBank.jsx'


export default class Payment extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            fromCard: true
        }
    }

    render() {
        return (
            <div>
                <button onClick={() => this.setState({fromCard: true})}>С карты любого банка</button>
                <button onClick={() => this.setState({fromCard: false})}>Из своего интернет-банка</button>
                {this.state.fromCard
                    ? <PaymentFromCard />
                    : <PaymentViaBank />}
            </div>
        )
    }
}