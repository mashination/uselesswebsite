import { publishReply } from '../store/actions/forum';
import { useDispatch, useSelector } from 'react-redux';
import { useForm } from "react-hook-form"

export default function ReplyForm(props){
    const dispatch = useDispatch();
    const { register, handleSubmit } = useForm();
    const onSubmit = data => console.log(data);
    const submitTopicHandler =async(data) => {
        console.log("submitti")
        await dispatch( publishReply (data.username,props.Id,data.content));
        props.setIsReply(false)
     }
    
    return (
        <div className='TFormContainter'>
            
           
            <form className='TForm' onSubmit={handleSubmit(submitTopicHandler)}>
            
                <input {...register("username",{ required: true})} placeholder="username" className='TFormUsr'/>
                <textarea {...register("content",{ required: true})} placeholder="content"className='TFormContent' />

                <input type="submit" className='TFormSubmit' />
            </form>
        </div>
    )
}