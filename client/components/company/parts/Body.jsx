import React from 'react';
import Request from '../payments/Request.jsx'
import Payment from '../payments/Payment.jsx'


export default class Body extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            isRequest: false
        }
    }

    render() {
        return (
            <div>
                <button onClick={() => this.setState({isRequest: false})}>Заплатить</button>
                <button onClick={() => this.setState({isRequest: true})}>Запросить платёж</button>
                {this.state.isRequest
                    ? <Request />
                    : <Payment />}
            </div>
        )
    }


}