/* eslint-disable no-unused-vars */
import React, { useEffect, useState } from 'react';
import { CardProduct, Navbar } from '../components';
import axios from 'axios';

const Products = () => {
  const [products, setProducts] = useState([]);
  const [searchTerm, setSearchTerm] = useState('');

  useEffect(() => {
    const fetchProducts = async () => {
      try {
        const response = await axios.get('http://localhost:3541/products');
        console.log(response.data);
        setProducts(response.data);
      } catch (error) {
        console.error('Error fetching products:', error);
      }
    };

    fetchProducts();
  }, []);
  const filteredProducts = products.filter((product) =>
    product.name.toLowerCase().includes(searchTerm.toLowerCase())
  );
  return (
    <div>
      <Navbar />
      <div className="flex justify-between items-start w-full h-24 bg-slate-300 px-10 flex-col">
        <div className="flex-grow">
          <input
            type="text"
            placeholder="Cari produk..."
            value={searchTerm}
            onChange={(e) => setSearchTerm(e.target.value)}
            className="px-4 py-2 border border-gray-300 rounded-md mr-4 my-2 w-full"
          />
        </div>
        <div className="flex-grow">
          <h1 className="text-md font-semibold">
            Berikut adalah beberapa produk dari para petani sekitar:
          </h1>
        </div>
      </div>

      <div className="flex flex-wrap justify-around">
        {filteredProducts.length > 0 ? (
          filteredProducts.map((product, index) => (
            <CardProduct key={index} product={product} />
          ))
        ) : (
          <p className="text-center mt-4">
            Tidak ada produk yang sesuai dengan pencarian.
          </p>
        )}
      </div>
    </div>
  );
};

export default Products;
