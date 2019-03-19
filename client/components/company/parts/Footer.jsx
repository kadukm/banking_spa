import React from 'react';
import {
    FacebookShareButton,
    GooglePlusShareButton,
    TwitterShareButton,
    TelegramShareButton,
    WhatsappShareButton,
    FacebookIcon,
    TwitterIcon,
    GooglePlusIcon,
    TelegramIcon,
    WhatsappIcon
} from "react-share";
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
                <div>
                    <div className="share">
                        <div>Поделитесь информацией о нас в соцсетях:</div>
                        <div className="share__icons">
                            <span className="share__icon">
                                <FacebookIcon size={24} round />
                            </span>
                            <span className="share__icon">
                                <TwitterIcon size={24} round />
                            </span>
                            <span className="share__icon">
                                <TelegramIcon size={24} round />
                            </span>
                            <span className="share__icon">
                                <WhatsappIcon size={24} round />
                            </span>
                            <span className="share__icon">
                                <GooglePlusIcon size={24} round />
                            </span>
                        </div>
                    </div>
                </div>
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
        