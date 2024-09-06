import React, { useContext, useState } from 'react';
import { ShopContext } from '../context/ShopContext';



const Login = () => {
  const [formState, setFormState] = useState('login');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [name, setName] = useState('');

  const { signIn, signUp ,user} = useContext(ShopContext);


  const handleSubmit = async (event) => {
    event.preventDefault();
    if (formState === 'login') {
      await signIn(email, password);
    } else {
      await signUp( email, password);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="flex flex-col items-center w-[90%] sm:max-w-96 m-auto mt-14 gap-4 text-gray-800">
      <div className="inline-flex items-center gap-2 mb-2 mt-10">
        <p className="prata-regular text-3xl">{formState === 'login' ? 'Login' : 'Sign Up'}</p>
        <hr className="border-none h-[1.5px] w-8 bg-gray-800" />
      </div>
      {formState === 'signUp' && (
        <input
          required
          type="text"
          value={name}
          onChange={(e) => setName(e.target.value)}
          className="w-full px-3 py-2 border border-gray-800"
          placeholder="Name"
        />
      )}
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
      <div className="w-full flex justify-between text-sm mt-[-8px]">
        <p className="cursor-pointer">Forget your password?</p>
        {formState === 'login' ? (
          <p onClick={() => setFormState('signUp')} className="cursor-pointer">
            Create account
          </p>
        ) : (
          <p onClick={() => setFormState('login')} className="cursor-pointer">
            Login Here
          </p>
        )}
      </div>
      <button type="submit" className="bg-black text-white font-light px-8 py-2 mt-4">
        {formState === 'login' ? 'Sign In' : 'Sign Up'}
      </button>
    </form>
  );
};

export default Login;
