import { Section, Cell, List, Button} from '@telegram-apps/telegram-ui';
import type { FC } from 'react';
import { useState, useEffect } from 'react';
import axios from 'axios';

//import { Link } from '@/components/Link/Link.tsx';
import { Pagination } from '@/components/Pagination/Pagination.tsx';
import { Page } from '@/components/Page.tsx';
import {Product, Card, ProductPhoto} from "@/components/Product/Card";

const httpClient = axios.create({
  baseURL: 'http://localhost/api',//import.meta.env.BASE_URL,
  headers: {
    'Content-Type': 'application/json'
  }
});

const detailUrl = function( ID :number) { return `/product/${ID}`; };

interface Paginator {
  limit: number;
  offset: number;
  total: number;
  page: number;
}

interface ProductsResponse {
  products: Product[];
  photos: ProductPhoto;
  paginator: Paginator;
}

export const IndexPage: FC = () => {
  const [products, setProducts] = useState<Product[]>();
  const [photos, setPhotos] = useState<ProductPhoto>();
  const [paginator, setPaginator] = useState<Paginator>({
    limit: 10,
    offset: 0, // ProductID as offset
    total: 0,
    page: 1
  });

  const getPhoto = (key: number): string[] => {
    let im : string[] = photos != null ? photos[key] : [];
    return im
  };

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
      setPhotos(productsResponse.photos);
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
        >
          {products ? (
            products.map(product => (
              <Card item={product} images={getPhoto(product.id)}/>
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