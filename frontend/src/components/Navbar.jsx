// eslint-disable-next-line no-unused-vars
import React from 'react';

const Navbar = () => {
  return (
    <nav className="bg-gray-800 p-4">
      <div className="container mx-auto flex justify-around items-center">
        <div className="font-bold">
          <a href="/" className="text-3xl text-yellow-400">
            Sayurku
          </a>
        </div>
        <div>
          <ul className="flex space-x-4">
            <li>
              <a href="" to="/" className="text-gray-300 hover:text-white">
                Home
              </a>
            </li>
            <li>
              <a href="/about"className="text-gray-300 hover:text-white">
                About
              </a>
            </li>
            <li>
              <a
                href="/contact"
                className="text-gray-300 hover:text-white"
              >
                Contact
              </a>
            </li>
            <li>
              <a
                href="/products"
                className="text-gray-300 hover:text-white"
              >
                Products
              </a>
            </li>
          </ul>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
