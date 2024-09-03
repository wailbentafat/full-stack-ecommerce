import React from 'react'
import Title from '../components/Title'
import { assets } from '../assets/assets'
import NewsLetterBox from '../components/NewsLetterBox'

const About = () => {
  return (
    <div>
      <div className='text-2xl text-center pt-8 border-t'>
        <Title text1={'ABOUT'} text2={'US'}/>
      </div>
      <div className='my-10 flex flex-cols md:flex-row gap-16'>
          <img className='w-full md:max-w-[450px]' src={assets.about_img} alt="" />
          <div className='flex flex-col justify-center gap-6 md:w-2/4 text-gray-600'>
              <p>Lorem ipsum dolor, sit amet consectetur adipisicing elit. Quasi, molestias maiores temporibus corporis consequatur quod cumque dignissimos, alias inventore magnam excepturi, ratione deleniti quo quia. Reiciendis excepturi eveniet, unde ex sapiente reprehenderit quia eum amet architecto provident expedita laboriosam soluta suscipit qui nisi nihil corporis quis autem dignissimos! Quo, quasi!</p>
              <p>Lorem ipsum dolor, sit amet consectetur adipisicing elit. Quasi, molestias maiores temporibus corporis consequatur quod cumque dignissimos, alias inventore magnam excepturi, ratione deleniti quo quia. Reiciendis excepturi eveniet, unde ex sapiente reprehenderit quia eum amet architecto provident expedita laboriosam soluta suscipit qui nisi nihil corporis quis autem dignissimos! Quo, quasi!</p>
              <b className='text-gray-800'>Our Mission</b>
              <p>Lorem ipsum dolor sit, amet consectetur adipisicing elit. Suscipit aspernatur explicabo eaque voluptatibus sint ea nisi voluptate. Fuga laudantium unde, a eaque alias aliquid quas adipisci vero, ducimus quisquam sit?</p>
          </div>
      </div>
      <div className='text-xl py-4'>
        <Title text1={'WHY'} text2={'CHOSE US'}/>
      </div>
      <div className='flex flex-col md:flex-row text-sm mb-20'>
        <div className='border px-10 md:px-16 py-8 sm:py-20 flex flex-col gap-5'>
            <b>Quality Assurance:</b>
            <p className='text-gray-600'>Lorem ipsum dolor sit amet consectetur adipisicing elit. Velit maiores dolor quasi dolorum ex ut?</p>
         </div>
         <div className='border px-10 md:px-16 py-8 sm:py-20 flex flex-col gap-5'>
            <b>Convenience:</b>
            <p className='text-gray-600'>Lorem ipsum dolor sit amet consectetur adipisicing elit. Velit maiores dolor quasi dolorum ex ut?</p>
         </div>
         <div className='border px-10 md:px-16 py-8 sm:py-20 flex flex-col gap-5'>
            <b>Exceptional Customer Service:</b>
            <p className='text-gray-600'>Lorem ipsum dolor sit amet consectetur adipisicing elit. Velit maiores dolor quasi dolorum ex ut?</p>
         </div>
     </div>
     <NewsLetterBox />
    </div>
  )
}

export default About
