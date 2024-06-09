// eslint-disable-next-line no-unused-vars
import { Routes, Route } from 'react-router-dom';
import { About, Contact, Homepage, NotFound, Products } from './pages';
const App = () => {
  return (
    <Routes>
      <Route path="/" element={<Homepage />} />
      <Route path="/about" element={<About />} />
      <Route path="/contact" element={<Contact />} />
      <Route path="/products" element={<Products />} />
      <Route path="*" element={<NotFound />} />
    </Routes>
  );
};

export default App;
