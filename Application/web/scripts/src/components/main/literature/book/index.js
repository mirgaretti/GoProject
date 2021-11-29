

const Book = (props) => {
    const { name, link, cost, costWithDiscount, imgLink } = props;
    
    return (
        <li><b>{name}</b><br />{link}<br />
            Стоимость: {cost}<br />
            Стоимость со скидкой: {costWithDiscount}<br />
            <img src={imgLink} className="" alt=""/>
        </li>
    )
}

export default Book;