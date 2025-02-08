import { Section, Cell, Image, List, Button} from '@telegram-apps/telegram-ui';
import type { FC } from 'react';
import { useState, useEffect } from 'react';
import axios from 'axios';

import { Link } from '@/components/Link/Link.tsx';
import { Pagination } from '@/components/Pagination/Pagination.tsx';
import { Page } from '@/components/Page.tsx';

import productSvg from './product.svg';

const httpClient = axios.create({
  baseURL: 'http://localhost/api',
  headers: {
    'Content-Type': 'application/json'
  }
});

const detailUrl = function( ID :number) { return `/product/${ID}`; };

interface Product {
  id: number;
  name: string;
  price: string;
  description: string;
}

interface Paginator {
  limit: number;
  offset: number;
  total: number;
  page: number;
}

interface ProductsResponse {
  products: Product[];
  paginator: Paginator;
}

export const IndexPage: FC = () => {
  const [products, setProducts] = useState<Product[]>();
  const [paginator, setPaginator] = useState<Paginator>({
    limit: 10,
    offset: 0, // ProductID as offset
    total: 0,
    page: 1
  });

  useEffect(() => {
    fetchProducts()
  }, []);

  const fetchProducts = () => {
    const offset = products? products[products.length - 1]?.id : 0;
    const listUrl = '/products?limit='+paginator.limit+'&offset='+offset;
    httpClient.get(listUrl).then(response => {
      const productsResponse: ProductsResponse = response.data as ProductsResponse;
      setProducts(productsResponse.products);
      setPaginator(productsResponse.paginator);
    }).catch(err => {
      console.log(err);
    });
  };

  return (
    <Page back={false}>
      <List>
        <Section
          header="Товары"
          footer="Новопокровское"
        >{products ? (
            products.map(product => (
              <Link to={detailUrl(product.id)} key={product.id}><Cell
                before={<Image src={productSvg} style={{ backgroundColor: '#464646' }}/>}
                subtitle={product.name}
              >
              <h1>{product.description}</h1>
              <p>Цена: {product.price}</p>
              </Cell>
              </Link>
            ))
          ) : (
            <Cell>Загрузка данных...</Cell>
          )}
          <Pagination paginator={paginator} />
          <Button onClick={fetchProducts}>Обновить данные</Button>
        </Section>
      </List>
    </Page>
  );
};