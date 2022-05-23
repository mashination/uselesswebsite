import React, {useEffect, useState, useCallback}from 'react';
import * as forumActions from '../store/actions/forum';
import { Squares } from "react-activity";
import { useDispatch, useSelector } from 'react-redux';
import { Outlet, Link } from "react-router-dom";
import '../App.css'


export default function TopicList(props){
    const topics = useSelector(state => state.forum.topics);
    const [isLoading, setIsLoading] = useState(true);
    const dispatch = useDispatch();

   const loadTopics = useCallback(async ()=> {
        try {
            await dispatch(forumActions.getTopics())
        } catch (error) {
            console.log(error.message)
        }
    })

    useEffect(()=>{
        setIsLoading(true)
        loadTopics().then(() => setIsLoading(false))
        console.log(topics)
        
    }, []);
    if (isLoading) {
        return (
            <div className='TList'>
                caca boudin de merde
                <Squares size={10000} color='white' />
                caca prout
            </div>
        )
    }
    return (
        <div className='ForumPage'>
            <Link  className='ForumLink'  to="/forum/newtopic">new topic</Link>
            {/* <h2 style={{ color: "white" }}>topic list</h2> */}
            <div className='TList' style={{ display: "flex" }}>
                {topics ? 
                <nav
                    style={{
                    borderRight: "solid 1px",
                    padding: "1rem",
                    }}
                >
                    {topics.map((topic) => (
                        <div className='TLinkContainer'>
                    <Link
                        onClick={()=>{props.setTlist(false); dispatch(forumActions,forumActions.clearCurrentTopic())}}
                        style={{  color: "white" }}
                        to={`/forum/${topic.Id}`}
                        key={topic.Id}
                        className='TLink'
                    >
                         <p className='TLinkTitle'>{topic.Title}</p>
                        <p className='TLinkUsr'>{topic.Usr}</p>
                    </Link>
                    </div>
                    ))}
                </nav> : <div color='white'>no topics</div>}
                <Outlet />
                </div>
        </div>
    )
}