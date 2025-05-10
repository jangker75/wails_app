import { useState } from 'react';
import './App.css';
import { SelectAndImportExcel } from "../wailsjs/go/backend/importer"; // Import fungsi dari Wails
import { backend } from '../wailsjs/go/models';
import { Button } from './components/ui/button';

function App() {
    const [data, setData] = useState<backend.ExcelData>({
        filename: "",
        header: [],
        details: []
    });
    const [error, setError] = useState<string | null>(null);

    const handleImport = async () => {
        try {
            setError(null);
            const result = await SelectAndImportExcel();
            // console.log("result", result);

            setData(result);
        } catch (err: any) {
            console.error(err);
            setError("Gagal mengimpor file Excel.");
        }
    };
    const handleDelete = async () => {
        try {
            console.log("handleDelete clicked");
            
        }
        catch (err: any) {
            console.error(err);
            setError("Gagal menghapus data.");
        }
    }

    return (
        <div className="container w-screen mx-auto my-10">
            <h1>Import File Excel</h1>
            <Button className='bg-blue-500 text-white hover:cursor-pointer' onClick={handleImport}>Pilih File Excel</Button>
            <Button className='bg-red-500 text-white hover:cursor-pointer' onClick={handleDelete}>Delete Data</Button>

            {error && <p style={{ color: "red" }}>{error}</p>}

            {
                data.details.length === 0 && <li>Tidak ada data yang diimpor.</li>
            }
            {data.details.length > 0 && (
                <>
                    <div className='my-3'><p className='font-bold'>Data yang diimpor:</p> {data.filename}</div>
                    <div className='flex justify-center'>
                        <table className='table-auto border border-gray-400 '>
                            <thead>
                                <tr>
                                    {data.header.map((item: any, idx: number) => (
                                        <th className='border border-gray-400' key={idx}>{item}</th>
                                    ))}

                                </tr>
                            </thead>
                            <tbody>
                                {data.details.map((item: string[], idx: number) => (
                                    <tr key={idx}>
                                        {item.map((subItem: string, subIdx: number) => (
                                            <td className='border border-gray-400 px-4' key={subIdx}>{subItem}</td>
                                        ))}
                                    </tr>
                                ))}

                            </tbody>
                        </table>
                    </div>

                </>
            )}

        </div>
    );
}

export default App
