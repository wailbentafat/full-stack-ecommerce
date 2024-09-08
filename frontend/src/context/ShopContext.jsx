import { createContext, useEffect, useState } from "react";
import { products } from "../assets/assets";
import { toast } from "react-toastify";
import { useNavigate } from "react-router-dom"

 export const ShopContext = createContext();

 const ShopContextProvider = (props) =>{

    const currency = ' DA'
    const [delivery_fee , setDelivery_fee] = useState(0);
    const [search ,setSearch]=useState('');
    const[showSearch ,setShowSearch] = useState(false);
    const [cartItems,setCartItems] = useState({})
    const navigate = useNavigate()
   

    const addToCart = async (itemId,size) => {

            if (!size) {
                toast.error('Select Product Size');
                return ;
            }


        let cartData = structuredClone(cartItems);
        if (cartData[itemId]) {
            if (cartData[itemId][size]) {
                cartData[itemId][size] +=1 ;
            }
            else{
                cartData[itemId][size] = 1;
            }
        }
        else{
            cartData[itemId] ={};
            cartData[itemId][size] =1 ;
        }
        setCartItems(cartData)
    }

    const getCartCount = () =>{
        let totaleCount = 0 ;
        for (const items in cartItems) {
           for(const item in cartItems[items]){
                try {
                    if (cartItems[items][item] > 0) {
                        totaleCount += cartItems[items][item]
                    }
                } catch (error) {
                    
                }
           }
        }
        return totaleCount ;
    }
  

    const updateQuantity = async (itemId ,size ,quantity) => {
        let cartData = structuredClone(cartItems);
        cartData[itemId][size] = quantity;

        setCartItems(cartData);
    }
 
    const getCartAmount = () => {
        let totalAmount = 0 ;
        for(const items in cartItems){
            let itemInfo = products.find((product)=> product._id === items);
            for(const item in cartItems[items]){
                try {
                    if (cartItems[items][item] > 0) {
                        totalAmount += itemInfo.price * cartItems[items][item];
                    }
                } catch (error) {
                    
                }
            }
        }
        return totalAmount ;
    }
    const [user, setUser] = useState(null);

    const signIn = async (email, password) => {
      try {
        // Call API to sign in
        const response = await fetch('http://localhost:8080/login', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ email, password }),
        });
        if(!response.ok) throw new Error('Sign in failed')
        const data = await response.json();
        Cookies.set('authToken',data.token,{expires:10})
        setUser(data.user);
      } catch (error) {
        console.error(error);
      }
    };
  
    const signUp = async ( email, password) => {
      try {
        // Call API to sign up
        const response = await fetch('http://localhost:8080/register', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ email, password }),
        });
        const data = await response.json();
        setUser(data.user);
        } catch (error) {
          console.error(error);
          }
          }

    const value ={
        products, currency ,delivery_fee,setDelivery_fee,
        search ,setSearch ,showSearch ,setShowSearch,cartItems,
        addToCart,getCartCount,updateQuantity,getCartAmount,navigate,user,signIn,signUp
    }
    return(
        <ShopContext.Provider value={value}>
            {props.children}
        </ShopContext.Provider>
    )
 }

 export default ShopContextProvider;
