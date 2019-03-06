import React from 'react';
import apiBaseUrl from '../../../config.js' 

export default class Header extends React.Component {
    constructor(props) {
        super(props)
        this.state = {info: undefined}
    }

    componentDidMount() {
        fetch(`${apiBaseUrl}/api/companies/${this.props.companyID}`)
            .then(response => response.json())
            .then(res => this.setState({info: res}))
    }

    render() {
        return (
            this.state.info !== undefined
                ? this.renderHeader()
                : null
        )
    }

    renderHeader = () => {
        return (
            <section>
                <img src={this.state.info.photo_path} alt="user photo" />
                <h3>{this.state.info.status} {this.state.info.name}</h3>
                <div>
                    <span>{this.state.info.phone} </span>
                    <a href={"http://" + this.state.info.site}>{this.state.info.site}</a>
                    <a href={"mailto:" + this.state.info.email}>{this.state.info.email}</a>
                </div>
                <a href={this.state.info.full_info_path}>Информация о компании</a>
                <a href={this.state.info.requisites_path}>Показать реквизиты</a>
            </section>
        )
    }
}