import { useState } from 'react';
import './App.css';
import { DeleteAllDataFromDB, SelectAndImportExcel } from "../wailsjs/go/backend/importer"; // Import fungsi dari Wails
import { backend, models } from '../wailsjs/go/models';
import { Button } from './components/ui/button';
import { toast } from 'sonner';
// import { toast } from 'sonner';
const initialDataImport: backend.ExcelData = {
    isSaveDB: false,
    filename: "",
    header: [],
    details: []
}
function App() {
    const [data, setData] = useState<backend.ExcelData>(initialDataImport);
    const [error, setError] = useState<string | null>(null);

    const handleImport = async () => {
        try {
            setError(null);

            const result: backend.ExcelData = await SelectAndImportExcel();
            if (result.details.length === 0) {
                toast.error("Tidak ada data yang diimpor.");
                return
            }
            setData(result);
            if (result.isSaveDB) {
                toast.success("Import Success ", {
                    description: "Data disimpan ke database",
                });
            } else {
                toast.success("Import Success", {
                    description: "Data tidak disimpan ke database",
                });
            }
        } catch (err: any) {
            console.error(err);
            if (err !== "shellItem is nil") {
                toast.error("Gagal mengimpor file Excel");
                setError("Gagal mengimpor file Excel.");
            }
        }
    };
    const handleDelete = async () => {
        try {
            const res: models.Response = await DeleteAllDataFromDB()
            if (res.status === "success") {
                toast.success(res.message);
                setData(initialDataImport)
            } else {
                toast.error(res.message);
            }
        }
        catch (err: any) {
            console.error(err);
            setError("Gagal menghapus data.");
        }
    }

    return (
        <div className="container w-screen mx-auto my-10">
            <h1>Import File Excel</h1>
            <div className='flex justify-center gap-2 my-5'>
                <Button className='bg-blue-500 text-white hover:cursor-pointer' onClick={handleImport}>Pilih File Excel</Button>
                <Button className='bg-red-500 text-white hover:cursor-pointer' onClick={handleDelete}>Delete Data</Button>
            </div>

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
