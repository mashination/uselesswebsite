import { publishTopic } from '../store/actions/forum';
import { testTopic } from '../store/actions/forum';
import { useDispatch, useSelector } from 'react-redux';
import { useForm } from "react-hook-form"
import '../App.css'
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';

export default function TopicForm(){
    let navigate = useNavigate()
    const dispatch = useDispatch();
    const { register, handleSubmit } = useForm();

    const submitTopicHandler =async(data) => {
        console.log("submitti")
        await dispatch( publishTopic (data.username,data.title,data.content));
        navigate("/forum/all");
     }
  

    return (
        <div className='TFormContainter'>
            
            
            <p>error</p>
            <form className='TForm' onSubmit={handleSubmit(submitTopicHandler)}>
            
                <input {...register("username",{ required: true})} placeholder="username" className='TFormUsr'/>
                <input {...register("title",{ required: true})} placeholder="title" className='TFormTitle'/>
                <textarea {...register("content",{ required: true})} placeholder="content" className='TFormContent'/>

                <input type="submit" className='TFormSubmit'/>
            </form>
        </div>
    )
}