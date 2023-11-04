"use client";
import { useState } from "react";

const SellingPage = () => {
  const tabs = [
    { id: 0, label: "Semua" },
    { id: 1, label: "Sembako" },
    { id: 2, label: "Snack" },
    { id: 3, label: "Minuman" },
    { id: 4, label: "Disabled", disabled: true },
  ];

  const [activeTab, setActiveTab] = useState(0);

  const handleTabClick = (tabId) => {
    setActiveTab(tabId);
  };

  // Placeholder data for products (you can replace this with your actual data)
  const products = [
    { id: 1, name: "Product 1", price: 10, variants: ["Small", "Medium", "Large"] },
    { id: 2, name: "Product 2", price: 15, variants: ["Small", "Medium", "Large"] },
    { id: 3, name: "Product 3", price: 20, variants: ["Small", "Medium", "Large"] },
  ];

  const addToCart = (product, variant, quantity) => {
    // Implement your logic to add the product to the cart with the selected variant and quantity
    console.log(`Added ${quantity}x ${variant} of ${product.name} to the cart.`);
  };

  const [selectedVariant, setSelectedVariant] = useState("Small");
  const [selectedQuantity, setSelectedQuantity] = useState(1);

  return (
    <>
      <div className="h-[100vh] w-full flex gap-2">
        <div className="w-3/4 bg-red-100 min-h-full">
          <h1 className="text-lg font-semibold mb-2">Selling Page</h1>

          <div className="h-full bg-red-300">
            <ul className="flex flex-wrap -mb-px text-sm font-medium text-center text-gray-500 dark:text-gray-400">
              {tabs.map((tab) => (
                <li key={tab.id} className="mr-2">
                  <a
                    href="#"
                    onClick={() => handleTabClick(tab.id)}
                    className={`inline-flex items-center justify-center p-4 border-b-2 ${
                      tab.id === activeTab
                        ? "border-blue-600 text-blue-600 dark:text-blue-500 dark:border-blue-500"
                        : "border-transparent"
                    } rounded-t-lg ${
                      tab.disabled
                        ? "cursor-not-allowed text-gray-400 dark:text-gray-500"
                        : "hover:text-gray-600 hover:border-gray-300 dark:hover:text-gray-300"
                    } group`}
                  >
                    {tab.label}
                  </a>
                </li>
              ))}
            </ul>
            <div className="p-4">
              <h2 className="text-lg font-semibold">Products</h2>
              <div className="grid grid-cols-2 gap-4 mt-4">
                {products.map((product) => (
                  <div
                    key={product.id}
                    className="bg-white p-4 border rounded shadow"
                  >
                    <h3 className="text-xl font-semibold">{product.name}</h3>
                    <p className="text-gray-600">Price: ${product.price}</p>

                    {/* Variant selection */}
                    <div className="mt-2">
                      <label htmlFor="variant">Select Variant:</label>
                      <select
                        id="variant"
                        value={selectedVariant}
                        onChange={(e) => setSelectedVariant(e.target.value)}
                        className="p-2 border rounded"
                      >
                        {product.variants.map((variant) => (
                          <option key={variant} value={variant}>
                            {variant}
                          </option>
                        ))}
                      </select>
                    </div>

                    {/* Quantity selection */}
                    <div className="mt-2">
                      <label htmlFor="quantity">Select Quantity:</label>
                      <input
                        type="number"
                        id="quantity"
                        value={selectedQuantity}
                        onChange={(e) => setSelectedQuantity(e.target.value)}
                        className="p-2 border rounded"
                      />
                    </div>

                    <button
                      onClick={() =>
                        addToCart(product, selectedVariant, selectedQuantity)
                      }
                      className="mt-2 px-4 py-2 bg-blue-600 text-white rounded"
                    >
                      Add to Cart
                    </button>
                  </div>
                ))}
              </div>
            </div>
          </div>
        </div>
        <div className="w-1/4 bg-meta-6 min-h-full"></div>
      </div>
    </>
  );
};

export default SellingPage;
