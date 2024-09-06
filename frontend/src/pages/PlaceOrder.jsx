import React, { useContext, useState } from 'react'
import Title from '../components/Title'
import CartTotale from '../components/CartTotale'
import { assets } from '../assets/assets'
import { ShopContext } from '../context/ShopContext'

const PlaceOrder = () => {

  const [method,setMethod] = useState('cod');
  const { navigate ,setDelivery_fee ,delivery_fee} = useContext(ShopContext);

  return (
    <div className='flex flex-col sm:flex-row justify-between gap-4 pt-5 sm:pt-14 min-h-[80vh]  border-t'>
      {/*   left side */}
      <div className='flex flex-col gap-4 w-full sm:max-w-[480px] '>
          <div className='text-xl sm:text-2xl my-3'>
            <Title  text1={'DELIVERY'} text2={'INFORMATION'}/>
          </div>
          <div className='flex gap-3'>
            <input className='border border-gray-300 rounded py-1.5 px-3.5 w-full' type="text" placeholder='First name' />
            <input className='border border-gray-300 rounded py-1.5 px-3.5 w-full' type="text" placeholder='Last name' />
          </div>
          <input className='border border-gray-300 rounded py-1.5 px-3.5 w-full' type="email" placeholder='Email address' />
          <input className='border border-gray-300 rounded py-1.5 px-3.5 w-full' type="text" placeholder='Street name' />
          <div className='flex gap-3'>
            <input className='border border-gray-300 rounded py-1.5 px-3.5 w-full' type="text" placeholder='City' />
            <select  className='border border-gray-300 rounded py-1.5 px-3.5 w-full' value={delivery_fee} onChange={e =>(setDelivery_fee(e.target.value))}  name='liver' required>
            <option value='0'>State</option>
             <option value="950.00">ADRAR</option>
             <option value="700.00"> CHLEF</option>
             <option value="800.00">LAGHOUAT</option>
             <option value="700.00">OUM EL BOUAGHI</option>
             <option value="700.00"> BATNA</option>
             <option value="700.00">BEJAIA</option>
             <option value="800.00"> BISKRA</option>
             <option value="800.00"> BECHAR</option>
             <option value="500.00"> BLIDA</option>
             <option  value="700.00" >BOUIRA</option>
             <option  value="1300.00" > TAMANRASSET</option>
             <option  value="700.00" > TEBESSA</option>
             <option  value="800.00" >TLEMCEN</option>
             <option  value="700.00" > TIARET</option>
             <option  value="600.00" > TIZI OUZOU</option>
             <option  value="400.00" > ALGER</option>
             <option  value="800.00" >DJELFA</option>
             <option  value="700.00" > JIJEL</option>
             <option  value="700.00" >SETIF</option>
             <option  value="700.00" >SAIDA</option>
             <option  value="700.00" > SKIKDA</option>
             <option  value="700.00" >SIDI BEL ABBES</option>
             <option  value="700.00" > ANNABA</option>
             <option  value="700.00" >GUELMA</option>
             <option  value="700.00" >CONSTANTINE</option>
             <option  value="650.00" >MEDEA</option>
             <option  value="700.00" >MOSTAGANEM</option>
             <option  value="750.00" >M SILA</option>
             <option  value="700.00" > MASCARA</option>
             <option  value="800.00" > OUARGLA</option>
             <option  value="600.00" >ORAN</option>
             <option  value="850.00">EL BAYADH</option>
             <option  value="0.00" > ILLIZI</option>
             <option  value="700.00" >BORDJ BOU ARRERIDJ</option>
             <option  value="550.00" >BOUMERDES</option>
             <option  value="700.00" >EL TARF</option>
             <option  value="0.00"> TINDOUF</option>
             <option  value="750.00" > TISSEMSILT</option>
             <option  value="800.00" >EL OUED</option>
             <option  value="700.00" >KHENCHELA</option>
             <option  value="700.00" >SOUK AHRAS</option>
             <option  value="600.00" > TIPAZA</option>
             <option  value="700.00" > MILA</option>
             <option  value="700.00" >AIN DEFLA</option>
             <option  value="850.00" >NAAMA</option>
             <option  value="700.00" > AIN TEMOUCHENT</option>
             <option  value="800.00" > GHARDAIA</option>
             <option  value="700.00" >RELIZANE</option>
             <option  value="1000.00"> TIMIMOUN</option>
             <option  value="0.00" > BORDJ BADJI MOKHTAR</option>
             <option  value="800.00" > OULED DJELLAL</option>
             <option  value="850.00" > BENI ABBES</option>
             <option  value="1300.00" >IN SALAH</option>
             <option  value="1300.00">IN GUEZZAM</option>
             <option  value="800.00">TOUGGOURT</option>
             <option  value="" >DJANET</option>
             <option  value="800.00" > EL M GHAIR</option>
             <option  value="850.00" >EL MENIAA</option>
          </select>
          </div>
          <div className='flex gap-3'>
            <input className='border border-gray-300 rounded py-1.5 px-3.5 w-full' type="text" placeholder='Zipcode' />
            <input className='border border-gray-300 rounded py-1.5 px-3.5 w-full' type="text"  value={'ALGERIA'} />
          </div>
          <input className='border border-gray-300 rounded py-1.5 px-3.5 w-full' type="number" placeholder='Phone'  />
          
      </div>

      {/*    right side */}
      <div className='mt-8'>
          <div className='mt-8 min-w-80'>
              <CartTotale />
          </div>
          <div className='mt-12'>
              <Title text1={'PAYMENT'} text2={'METHOD'}/>
              {/*   payement method section */}
              <div className='flex  gap-3 flex-col lg:flex-row'>
                  <div onClick={()=>setMethod('stripe')} className='flex items-center gap-3 border p-2 px-3 cursor-pointer'>
                    <p className={`min-w-3.5 h-3.5 border rounded-full ${method ==='stripe' ? 'bg-green-400' : ''}`}></p>
                    <img  className='h-5 mx-4' src={assets.stripe_logo}   alt="" />
                  </div>
                  <div   onClick={()=>setMethod('razorpay')} className='flex items-center gap-3 border p-2 px-3 cursor-pointer'>
                    <p className={`min-w-3.5 h-3.5 border rounded-full ${method ==='razorpay' ? 'bg-green-400' : ''}`}></p>
                    <img  className='h-5 mx-4' src={assets.razorpay_logo}   alt="" />
                  </div>
                  <div  onClick={()=>setMethod('cod')} className='flex items-center gap-3 border p-2 px-3 cursor-pointer'>
                    <p className={`min-w-3.5 h-3.5 border rounded-full ${method ==='cod' ? 'bg-green-400' : ''}`}></p>
                    <p className='text-gray-500 text-sm font-medium mx-4'>CASH ON DELIVERY</p>
                  </div>
              </div>

              <div className='w-full text-end mt-8'>
                    <button onClick={()=>navigate('/orders')} className='bg-black text-white px-16 py-3 text-sm'>PLACE ORDER</button>
              </div>



          </div>
      </div>
    </div>
  )
}

export default PlaceOrder
