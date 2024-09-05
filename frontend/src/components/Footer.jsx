import React from 'react'
import { assets } from '../assets/assets'
import { Link } from 'react-router-dom'

const Footer = () => {
  return (
    <div>
      <div className='flex flex-col sm:grid grid-cols-[3fr_1fr_1fr] gap-14 my-10 mt-40 text-sm'>
        <div>
            <img src={assets.logo} className='mb-5 w-31' alt="" />
            <p className='w-full md:w-2/3 text-gray-600'>
                Lorem ipsum dolor sit amet consectetur, adipisicing elit. Corporis distinctio facere praesentium aliquid vel? Molestiae temporibus ullam neque tempora ratione.
            </p>
        </div>
        <div>
            <p className='text-xl font-medium mb-5'>COMPANY</p>
            <ul className='flex flex-col gap-1 text-gray-600'>
                    <li><Link to="/">Home</Link> </li>
                    <li><Link to="/about">About us </Link></li>
                    <li><Link to="/">Delivery</Link></li>
                    <li><Link to="/">Privacy policy</Link></li>
            </ul>
        </div>
        <div>
            <p className='text-xl font-medium mb-5'>GET IN TOOCH</p>
            <ul className='flex flex-col gap-1 text-gray-600'>
                <li>+213-657-13-79-28</li>
                <li>zakihamidi385@gmail.com</li>
            </ul>
        </div>
      </div>
      <div>
        <hr />
        <p className='py-5 text-sm text-center'>Copyright 2024@ forever.com - All Right Reserved.</p>
      </div>
    </div>
  )
}

export default Footer
