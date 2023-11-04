import Calendar from "@/components/Calender";
import { Metadata } from "next";

export const metadata: Metadata = {
  title: "Inventory Page",
  description: "Inventory Page Pos Mania Mantap",

  
  // other metadata
};

const InventoryPage = () => {
  return (
    <>
      <div className="h-[100vh] w-full flex gap-2">
        <div className="w-3/4 bg-red-100 min-h-full">
          <h1>
            Inventory Page
          </h1>
        </div>
        <div className="w-1/4 bg-meta-6 min-h-full"></div>

      </div>
    </>
  );
};

export default InventoryPage;
