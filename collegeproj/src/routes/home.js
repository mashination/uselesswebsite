import { useParams } from "react-router-dom";
import React, {useEffect, useState, useCallback}from 'react';
import * as forumActions from '../store/actions/forum';
import { Squares } from "react-activity";
import { useDispatch, useSelector } from 'react-redux';
import { Outlet, Link } from "react-router-dom";
import ReplyForm from "../components/replyForm";
import '../App.css';




export default function Home() {
  const launches = useSelector(state => state.forum.launches);
    const [isLoading, setIsLoading] = useState(true);
    const [isReload, setReload] = useState(true);
    const dispatch = useDispatch();

    const loadLaunches = useCallback(async ()=> {
        try {
            await dispatch(forumActions.getLaunches())
        } catch (error) {
            console.log(error.message)
        }
    })

    useEffect(()=>{
        setIsLoading(true)
        loadLaunches().then(() => {setIsLoading(false); setReload(false)})
        console.log(launches)
      
    }, [isReload]);
    if (isLoading) {
        return (
            <div>
                <Squares/>
            </div>
        )
    }
    return (
      <main style={{ padding: "1rem 0" }}>
        <div className="HomeTextContainer">
        <p className="HomeText">
          I had to make a website for a college project. The topic was free, however it had to fit some requirements.
          Therefore, as I have no creativity, I present to you The Useless Website that (hopefully) fits all of them.
        </p >
        <p className="HomeText">
          This website has to fit the following:
        </p>
        <p className="HomeText">
          Back end written in Go net/http module wihtout using a framework. This can be seen in the code.
        </p>
        <p className="HomeText"> 
          It has to include user generated content. For this one you may find an anonymous forum under the Forum menu.
        </p>
        <p className="HomeText">
          It has to use an external API with dynamic content. Below you may find a list of the 5 next rocket launches, because don't we all want to leave this earth at this point.
        </p >
        </div>
        {launches ?
                    <ul
                        style={{
                        borderRight: "solid 1px",
                        padding: "1rem",
                        }}
                    >
                        {launches.map((launch) => 
                            <li
                                className="HomeLaunchContainer"
                                style={{  color: "white" , display: "block", margin: "1rem 0" }}
                                key={launch.Id}
                                
                            >
                                <div className="HLaunchContainer">
                                    <p className="HomeLaunch">{launch.Message}</p>
                                </div>
                                
                            </li>
                                
                            
                        )}
                    </ul> 
                    : <div color="white">Loading ... </div>
                }
              <button className="RTopReply" style={{ color: "black" }} onClick={() => setReload(true)}>Reload</button>
      </main>
    );
  }