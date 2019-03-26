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
                <div className="navbar__second-row">
                    <button className={"navbar__choice-second" + (this.state.fromCard ? " button-selected" : "")}
                        onClick={() => this.setState({fromCard: true})}
                    >
                        💳 С карты любого банка
                    </button>
                    <button className={"navbar__choice-second" + (this.state.fromCard ? "" : " button-selected")}
                        onClick={() => this.setState({fromCard: false})}
                    >
                        💻 Из своего интернет-банка
                    </button>
                </div>
                {this.state.fromCard
                    ? <PaymentFromCard />
                    : <PaymentViaBank />}
            </div>
        )
    }
}