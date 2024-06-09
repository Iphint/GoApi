/* eslint-disable no-unused-vars */
/* eslint-disable react/prop-types */
// eslint-disable-next-line no-unused-vars
import axios from 'axios';
import React, { useEffect, useState } from 'react';

const CardProduct = ({ product }) => {
  const [user, setUser] = useState(null);

  useEffect(() => {
    const fetchUser = async () => {
      try {
        const response = await axios.get(
          `http://localhost:3541/users/${product.user_id}`
        );
        console.log(response.data);
        setUser(response.data);
      } catch (error) {
        console.error('Error fetching user:', error);
      }
    };

    fetchUser();
  }, [product.user_id]);
  return (
    <div className="max-w-sm rounded overflow-hidden shadow-lg bg-white m-3">
      <div
        className="w-30 h-60 bg-cover bg-center rounded-lg"
        style={{
          backgroundImage:
            'url(https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQHTn7ih7xjeTUpQqQs5R7gj6THB-2xXes5Tw&s)',
        }}
      />
      <div className="px-6 py-4">
        <div className="font-bold text-xl mb-2">{product.name}</div>
        <p className="text-gray-700 text-base">{product.condition}</p>
        <p className="text-gray-700 text-base">{product.price}</p>
        {user && (
          <p className="text-gray-700 text-base">Created by: {user.username}</p>
        )}
      </div>
      <div className="px-6 pt-4 pb-2">
        <span className="inline-block bg-gray-200 rounded-full px-3 py-1 text-sm font-semibold text-gray-700 mr-2">
          #tag1
        </span>
        <span className="inline-block bg-gray-200 rounded-full px-3 py-1 text-sm font-semibold text-gray-700 mr-2">
          #tag2
        </span>
        <span className="inline-block bg-gray-200 rounded-full px-3 py-1 text-sm font-semibold text-gray-700">
          #tag3
        </span>
      </div>
    </div>
  );
};

export default CardProduct;
