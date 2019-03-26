import React from 'react';
import Modal from 'react-modal'
import Request from '../payments/Request.jsx'
import Payment from '../payments/Payment.jsx'
import '../../../styles/Body.css'

Modal.setAppElement('#app')

export default class Body extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            isRequest: false
        }
    }

    render() {
        return (
            <section>
                <div className="navbar__first-row">
                    <button className={"navbar__choice" + (this.state.isRequest ? "" : " button-selected")}
                        onClick={() => this.setState({isRequest: false})}
                    >
                        Заплатить
                    </button>
                    <button className={"navbar__choice" + (this.state.isRequest ? " button-selected" : "")}
                        onClick={() => this.setState({isRequest: true})}
                    >
                        Запросить платёж
                    </button>
                </div>
                {this.state.isRequest
                    ? <Request status={this.props.status} name={this.props.name}/>
                    : <Payment />}
            </section>
        )
    }
}