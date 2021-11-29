import Book from "./book";

const Literature = () => {
    const books = [{ 
        name: 'Математика. ЦТ. Теория. Примеры. Тесты',
        link: 'https://oz.by/edubooks/more10241684.html',
        cost: 14.90,
        costWithDiscount: 11.18,
    },
    { 
        name: 'Математика. ЦТ. Теория. Примеры. Тесты',
        link: 'https://oz.by/edubooks/more10241684.html',
        cost: 14.90,
        costWithDiscount: 11.18,
    },
    { 
        name: 'Математика. ЦТ. Теория. Примеры. Тесты',
        link: 'https://oz.by/edubooks/more10241684.html',
        cost: 14.90,
        costWithDiscount: 11.18,
        imgLink: 'dasd',
    }];

    return (
        <div>
            <div class="list-description">
                <h1>
                    Предлагаемая литература
                </h1>
            </div>
            <div class="list-description">
                <ol class="pills">
                    {books.map((book) => 
                    <Book 
                        name={book.name} 
                        link={book.link}
                        cost={book.cost}
                        costWithDiscount={book.costWithDiscount}
                        imgLink={book?.imgLink}
                    />)}
                </ol>
            </div>
            <div class="stat-footer-img"></div>
            <div class="literature-back-img"></div>
        </div>
    );
}

export default Literature;