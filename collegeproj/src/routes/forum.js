import { useEffect, useState } from "react";

import { Link } from "react-router-dom";
import { Outlet } from "react-router-dom";
import { useDispatch, useSelector } from 'react-redux';
import '../App.css';


export default function Forum() {
  const [tlist, setTlist] = useState(true)
  const dispatch = useDispatch();
  

    return (
      <main style={{ padding: "1rem 0" }} className="ForumPage">
        
         <Outlet/>
      </main>
    );
  }