import React from 'react';
import "../../../styles/Header.css"

export default class Header extends React.Component {
    render() {
        return (
            <section className="user-info box-wrapper">
                <img className="user-info__photo" src={this.props.photo_path} alt="user" />
                <h3 className="user-info__header">{this.props.status} {this.props.name}</h3>
                <div className="user-info__contacts">
                    <span className="user-info__phone">{this.props.phone} </span>
                    <a className="user-info__site" href={"http://" + this.props.site}>{this.props.site}</a>
                    <a href={"mailto:" + this.props.email}>{this.props.email}</a>
                </div>
                <a className="user-info__company" href={this.props.full_info_path}>Информация о компании</a>
                <a className="user-info__requisites" href={this.props.requisites_path}>Показать реквизиты</a>
            </section>
        )
    }
}