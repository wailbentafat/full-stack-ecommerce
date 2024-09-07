
import React from 'react';
import { Navigate, useLocation } from 'react-router-dom';
import { useAuth } from './authcontext'; 

const ProtectedRoute = ({ element, ...rest }) => {
    const { auth, loading } = useAuth();
    const location = useLocation();

    console.log('auth')
    return auth ? element : <Navigate to="/login" state={{ from: location }} replace />;
};

export default ProtectedRoute;
