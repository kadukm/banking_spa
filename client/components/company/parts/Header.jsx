import React from 'react';

export default class Header extends React.Component {
    render() {
        return (
            <section>
                <img src={this.props.photo_path} alt="user photo" />
                <h3>{this.props.status} {this.props.name}</h3>
                <div>
                    <span>{this.props.phone} </span>
                    <a href={"http://" + this.props.site}>{this.props.site}</a>
                    <a href={"mailto:" + this.props.email}>{this.props.email}</a>
                </div>
                <a href={this.props.full_info_path}>Информация о компании</a>
                <a href={this.props.requisites_path}>Показать реквизиты</a>
            </section>
        )
    }
}