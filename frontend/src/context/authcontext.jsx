import React, { useEffect, useState, createContext, useContext } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";
import Cookies from "js-cookie";

const AuthContext = createContext();

export const useAuth = () => useContext(AuthContext);

const AuthContextProvider = ({ children }) => {
    const [auth, setAuth] = useState(false);
    const [loading, setLoading] = useState(true);
    const navigate = useNavigate();
    const token = Cookies.get("token");

    useEffect(() => {
        const checkAuth = async () => {
            if (!token) {
                navigate("/login");
                return;
            }

            try {
                const response = await axios.get("http://localhost:8080/authorized", {
                    headers: {
                        "Authorization": `Bearer ${token}`,
                    },
                });

                console.log(response.data);
                if (response.data.auth) {
                    setAuth(true);
                } else {
                    setAuth(false);
                    navigate("/login");
                }
            } catch (error) {
                console.error(error);
                setAuth(false);
                navigate("/login");
            } finally {
                setLoading(false);
            }
        };

        checkAuth();
    }, [navigate, token]);

  
    return (
        <AuthContext.Provider value={{ auth, loading }}>
            {!loading && children}
        </AuthContext.Provider>
    );
};

export { AuthContextProvider };
