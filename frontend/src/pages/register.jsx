import React, { useState } from 'react';
import Cookies from 'js-cookie';
import { Link, useNavigate } from 'react-router-dom';
import axios from 'axios';

const Register = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState(null);
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const handleSubmit = async (event) => {
    event.preventDefault();
    setLoading(true);
    setError(null);

    try {
      const response = await axios.post('http://localhost:8080/register', {
        email,
        password,
      }, {
        headers: {
          'Content-Type': 'application/json',
        },
      });
      console.log('Response Status:', response.status);
      console.log(response.data);
      const data = response.data;

      if (data.token) {
        
        Cookies.set('token', data.token);
        navigate('/');
      } else {
        
        setError(data.error || 'An error occurred');
        console.log(data.error);
      }
    } catch (error) {
     console.log(error)
      setError('An error occurred. Please try again later.');
    }

    setLoading(false);
  };

  return (
    <form onSubmit={handleSubmit} className="flex flex-col items-center w-[90%] sm:max-w-96 m-auto mt-14 gap-4 text-gray-800">
    <div className="inline-flex items-center gap-2 mb-2 mt-10">
      <p className="prata-regular text-3xl">Sign Up</p>
      <hr className="border-none h-[1.5px] w-8 bg-gray-800" />
    </div>
    <input
      required
      type="email"
      value={email}
      onChange={(e) => setEmail(e.target.value)}
      className="w-full px-3 py-2 border border-gray-800"
      placeholder="Email"
    />
    <input
      required
      type="password"
      value={password}
      onChange={(e) => setPassword(e.target.value)}
      className="w-full px-3 py-2 border border-gray-800"
      placeholder="Password"
      minLength={8}
    />
    <div className="w-full flex  justify-end text-sm mt-[-8px]">
     <Link to={'/login'}><p   className="cursor-pointer ">
            Login Here 
        </p></Link> 
       
    </div>
    <button  type="submit" className="bg-black text-white font-light px-8 py-2 mt-4" disabled={loading}>{loading ? 'Registering...' : 'Register'}</button>
      <div className="flex items-center justify-between">
          
            {error && <p className="text-red-500">{error}</p>}
          </div>
          </form>
  );
};
export default Register;

