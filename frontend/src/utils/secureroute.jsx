import React, { useEffect, useState } from 'react';
import { Outlet, useNavigate, useLocation } from 'react-router-dom';
import cookies from 'js-cookie';

const SecureRoute = () => {
    const navigate = useNavigate();
    const location = useLocation();
    const [isAuthenticated, setIsAuthenticated] = useState(null); 

    useEffect(() => {
        const checkAuth = () => {
            const token = cookies.get('token');
            console.log('Token:', token);
            console.log('Current Path:', location.pathname);

            if (token) {
                setIsAuthenticated(true);
            } else {
                setIsAuthenticated(false);
            }
        };

        checkAuth();
    }, [location.pathname]); 

    useEffect(() => {
        if (isAuthenticated === false) {
            navigate('/login', { replace: true });
        }
    }, [isAuthenticated, navigate]);

    if (isAuthenticated === null) {
        return <div>Loading...</div>; 
    }

    return isAuthenticated ? <Outlet /> : null; 
};

export default SecureRoute;
