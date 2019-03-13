import React from 'react';
import apiBaseUrl from '../../../config.js'

export default class Request extends React.Component {
    postRequest = () => {
        console.log('request is sent')
    }

    render() {
        return (
            <div>
                <header>
                    <strong>
                        Создайте платежку, а {this.props.status} {this.props.name} подпишет её у себя в
                        интернет-банке
                    </strong>
                </header>
                <div>
                    <div>
                        <label htmlFor="inn">ИНН получателя</label>
                        <input type="text" name="inn" id="inn" required />
                    </div>
                    <div>
                        <label htmlFor="bik">БИК</label>
                        <input type="text" name="bik" id="bik" required />
                    </div>
                    <div>
                        <label htmlFor="account_number">Номер счёта</label>
                        <input type="text" name="account_number" id="account_number" required />
                    </div>
                    <div>
                        <label htmlFor="for_what">За что</label>
                        <input type="text" id="for_what" name="for_what" required />
                    </div>
                    <div>
                        <label htmlFor="amount">Сколько</label>
                        <input type="text" name="amount" id="amount" required />
                    </div>
                    <div>
                        <label htmlFor="phone">Телефон</label>
                        <input type="tel" name="phone" id="phone" placeholder="Ваш номер телефона" required/>
                    </div>
                    <div>
                        <label htmlFor="email">Ваш email</label>
                        <input type="email" name="email" id="email" required />
                    </div>
                    <button onClick={this.postRequest}>Заплатить</button>
                </div>
            </div>
        )
    }
}