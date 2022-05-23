import { useParams } from "react-router-dom";
import React, {useEffect, useState, useCallback}from 'react';
import * as forumActions from '../store/actions/forum';
import { Squares } from "react-activity";
import { useDispatch, useSelector } from 'react-redux';
import { Outlet, Link } from "react-router-dom";
import ReplyForm from "../components/replyForm";
import '../App.css';


export default function Topic(props) {
    let params = useParams();
    const topic = useSelector(state => state.forum.currentTopic);
    const [isLoading, setIsLoading] = useState(true);
    const [isReply, setIsReply] = useState(false);
    const dispatch = useDispatch();

    const loadTopic = useCallback(async ()=> {
        try {
            await dispatch(forumActions.getTopic(params.Id))
        } catch (error) {
            console.log(error.message)
        }
    })

    useEffect(()=>{
        setIsLoading(true)
        loadTopic().then(() => setIsLoading(false))
        console.log(topic)
      
    }, [isReply]);
    if (isLoading) {
        return (
            <div>
                <Squares/>
            </div>
        )
    }

    return (
        <div>
            <div className="RTopContainer">
                <div className="RTitleContainer">
                <p className="RTopUsr">{topic.Usr}</p>
                <h2 className="RTopTitle" style={{ color: "white" }}>{topic.Title}</h2>
                </div>
            
                <p className="RTopContent" style={{ color: "white" }}>{topic.Content}</p>
            </div>
            <div className="RListContainer" style={{ display: "flex" }}>
                {isReply ? <ReplyForm setIsReply = {setIsReply} Id={params.Id}/> : 

                <button className="RTopReply" style={{ color: "black" }} onClick={() => setIsReply(true)}>Reply</button>}

                {topic.Replies ?
                    <ul
                        style={{
                        borderRight: "solid 1px",
                        padding: "1rem",
                        }}
                    >
                        {topic.Replies.map((reply) => 
                            <li
                                className="RRepContainer"
                                style={{  color: "white" , display: "block", margin: "1rem 0" }}
                                key={reply.Id}
                                
                            >
                                <div className="RReplyContainer">
                                    <p className="RReplyUsr">{reply.Usr}</p>
                                    <h3 className="RReplyContent" style={{ color: "white" }}>{reply.Content}</h3>
                                </div>
                                
                            </li>
                                
                            
                        )}
                    </ul> : <div color="white">No replies</div>
                }
                <Outlet />
                </div>
        </div>
    )

}