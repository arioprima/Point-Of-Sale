import TableThree from "@/components/Tables/TableThree";
import { Metadata } from "next";

export const metadata: Metadata = {
  title: "Inventory Page",
  description: "Inventory Page Pos Mania Mantap",


  // other metadata
};

const InventoryPage = () => {
  return (
    <>
      <div className="h-[100vh] w-full ">
          <h1 className="text-lg font-semibold mb-4">
            Inventory Page
          </h1>
          <TableThree />

      </div>
    </>
  );
};

export default InventoryPage;
