import React from 'react';
import Modal from 'react-modal'

export default class Login extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            data: {
                login: "",
                password: ""
            },
            modal: {
                show: false,
                message: undefined
            }
        }
    }

    logIn = () => {
        fetch("/login", {method: "POST", body: JSON.stringify(this.state.data)})
            .then(response => response.json())
            .then(res => this.setState({modal: {show: true, message: res.result}}))
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
                <div>
                    <label htmlFor="login">Логин</label>
                    <input
                        type="text"
                        name="login"
                        id="login"
                        onChange={e => this.setState({data: {...this.state.data, login: e.target.value}})}
                        value={this.state.login}
                    />
                </div>
                <div>
                    <label htmlFor="password">Пароль</label>
                    <input
                        type="password"
                        name="password"
                        id="password"
                        onChange={e => this.setState({data: {...this.state.data, password: e.target.value}})}
                        value={this.state.password}
                    />
                </div>
                <button
                    onClick={this.logIn}
                >
                    Log in
                </button>
            </div>
        )
    }
}