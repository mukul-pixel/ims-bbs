import axios from 'axios';
import { toast, ToastContainer } from 'react-toastify';
import React, { useState } from 'react'
import 'react-toastify/dist/ReactToastify.css';

export const CreateGRN = () => {
    const [showModal, setShowModal] = useState(false);
    const [invoiceId,setInvoiceId] = useState()
    const [items, setItems] = useState([]);
    const [upc,setUpc] = useState("");
    const [quantity,setQuantity] = useState(0);
    const [cases , setCases] = useState(0);

    const handleUpcChange = async()=>{
      try {
        
        const response = await axios.get("http://localhost:8080/api/v1/findProductByUpc", {
            params: { upc } 
        });

        if (response.status === 200) {
            toast.success('Product added!');

            const currentUpc = upc;
            const currentCases = cases;
            const currentQuantity = quantity;

            const product = {
              Upc:currentUpc,
              Case:currentCases,
              Quantity:currentQuantity
            }

            setItems(prevItems=>[...prevItems,product])
        } else {
            toast.error('Unexpected response status: ' + response.status);
        }
    } catch (error) {
        if (error.response) {
            // Handle HTTP errors
            if (error.response.status === 404) {
                toast.error('Product not found!');
            } else if (error.response.status === 400) {
                toast.error('Bad request: ' + error.response.data.error);
            } else {
                toast.error('An error occurred:' + error.response.data.error);
            }
        } else if (error.request) {
            // Handle network errors
            toast.error('Network error. Please try again later.');
        } else {
            // Handle other errors
            toast.error('Error: ' + error.message);
        }
    }
  }

    const toggleModal = () => {
        setShowModal(prevState => !prevState);
    };
    
  return (
    <>
    <ToastContainer/>
    <div className='text-center py-3 bg-black text-white'>
        CREATE GRN
    </div>
    <div className='text-center py-3'>
        <div className='row py-2 invoiceDetails'>
            <span className='col-md-6'>
            <label htmlFor='invoiceId' className='form-label' style={{fontSize:"smaller"}}>InvoiceId:</label>
            <input type='text' id='invoiceId' className='rounded-2 ms-2'/>
            </span>
            <span className='col-md-6'>
            <label htmlFor='invoiceId' className='form-label' style={{fontSize:"smaller"}}>Merchant Name:</label>
            <input type='text' id='invoiceId' className='rounded-2 ms-2'/>
            </span>
        </div>
        <div className='row py-4 transportDetails'>
            <span className='col-md-4'>
            <label htmlFor='invoiceId' className='form-label' style={{fontSize:"smaller"}}>Transport Name</label>
            <input type='text' id='invoiceId' className='rounded-2 ms-2'/>
            </span>
            <span className='col-md-4'>
            <label htmlFor='invoiceId' className='form-label' style={{fontSize:"smaller"}}>Driver Name:</label>
            <input type='text' id='invoiceId' className='rounded-2 ms-2'/>
            </span>
            <span className='col-md-4'>
            <label htmlFor='invoiceId' className='form-label' style={{fontSize:"smaller"}}>Driver Contact:</label>
            <input type='text' id='invoiceId' className='rounded-2 ms-2'/>
            </span>
        </div>
    </div>
    <div className=''>
        <div>
        <table style={{ width: '96%', margin: '0 auto', fontSize:"smaller" }} className='text-center rounded-4'>
                <thead style={{ backgroundColor: '#f2f2f2' }} className='rounded-4'>
                    <tr>
                        <th style={{ padding: '7px' }}>S.no</th>
                        <th style={{ padding: '7px' }}>InvoiceId</th>
                        <th style={{ padding: '7px' }}>ItemId</th>
                        <th style={{ padding: '7px' }}>ItemUpc</th>
                        <th style={{ padding: '7px' }}>Total Cases</th>
                        <th style={{ padding: '7px' }}>Total Quantity</th>
                    </tr>
                </thead>
                <tbody>
                    {items.length === 0 ? (
                        <tr>
                            <td colSpan="6" style={{ padding: '7px', textAlign: 'center' }}>
                                No items available
                            </td>
                        </tr>
                    ) : (
                        items.map((item, index) => (
                            <tr className='GRNItemRows' key={item.id}>
                                <td style={{ padding: '7px' }}>{index + 1}</td>
                                <td style={{ padding: '7px' }}>{item.invoiceId}</td>
                                <td style={{ padding: '7px' }}>{item.itemId}</td>
                                <td style={{ padding: '7px' }}>{item.itemUpc}</td>
                                <td style={{ padding: '7px' }}>{item.totalCases}</td>
                                <td style={{ padding: '7px' }}>{item.totalQuantity}</td>
                            </tr>
                        ))
                    )}
                </tbody>
            </table>
            <div className='text-end m-3'>
                <button type='button' className='createGRNButtons me-2' onClick={toggleModal}>Add Item</button>
                {showModal && (
        <div className="modal fade show" tabIndex="-1" role="dialog" style={{ display: 'block' }}>
        <div className="modal-dialog modal-md modal-dialog-centered" role="document">
        <div className="modal-content modalSubscribe">
      <div className='row p-2'>
        <div className='text-center col-11'>Add GRN Item</div>
        <button type="button" className="btn-close col-1 bg-light me-2" onClick={toggleModal}></button>
      </div>
      <div className="modal-body text-white text-center">
        <form className='p-3'>
          <div className="mb-3">
            <input
              type="text"
              className="form-control mb-1"
              id="upc-input"
              placeholder="Product-Upc"
              onChange={(e)=>setUpc(e.target.value)}
            />
            <input
              type="number"
              className="form-control mb-1"
              id="cases-input"
              placeholder="Total-Cases"
              onChange={(e) => setCases(e.target.value)}
            />
            <input
              type="number"
              className="form-control"
              id="quantity-input"
              placeholder="Total-Quantity"
              onChange={(e) => setQuantity(e.target.value)}
            />
          </div>
        </form>
      </div>
      <div className='p-4'>
        <button type="button" className="btn btn-light form-control" onClick={handleUpcChange}>Submit</button>
      </div>
    </div>
        </div>
      </div>  
      )}
                <button type='submit' className='createGRNButtons me-2'>Create GRN</button>
            </div>
        </div>
    </div>
    </>
  )
}
