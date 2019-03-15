import React from 'react';
import "../../../styles/Footer.css"

export default class Footer extends React.Component {
    render() {
        return (
            <section>
                <h3>О компании {this.props.status} {this.props.name}</h3>
                <div className="products">
                    {this.props.products.map(product => <Product {...product} key={product.name}/>)}
                </div>
                <div className="products__info">
                    {this.props.info}
                </div>
                <a href={this.props.full_info_path}>Полная информация</a>
            </section>
        )
    }
}

function Product(props) {
    return (
        <div className="product">
            <img src={props.image_path} alt="product_image" />
            <div><strong>{props.name}</strong></div>
            <div>{props.price}</div>
        </div>
    )

}
        