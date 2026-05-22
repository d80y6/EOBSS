import React from 'react';

export default function InventoryDashboard() {
  return (
    <div className="p-8 space-y-6">
      <header className="flex justify-between items-center">
        <h1 className="text-2xl font-bold">Network & Service Inventory</h1>
        <div className="flex gap-2">
          <input
            type="text"
            placeholder="Search resources..."
            className="px-4 py-2 border rounded-md"
          />
        </div>
      </header>

      <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
        <InventoryCard title="Physical Resources" count="12,450" status="healthy" />
        <InventoryCard title="Logical Resources" count="45,800" status="warning" />
        <InventoryCard title="Active Services" count="1,150,000" status="healthy" />
      </div>

      <table className="w-full bg-white rounded-lg shadow border">
        <thead className="bg-slate-50">
          <tr>
            <th className="p-4 text-left">Resource ID</th>
            <th className="p-4 text-left">Name</th>
            <th className="p-4 text-left">Type</th>
            <th className="p-4 text-left">Status</th>
            <th className="p-4 text-left">Actions</th>
          </tr>
        </thead>
        <tbody>
          <ResourceRow id="RES-001" name="NYC-OLT-01" type="OLT" status="Active" />
          <ResourceRow id="RES-002" name="NYC-CORE-02" type="Router" status="Warning" />
          <ResourceRow id="RES-003" name="SIM-998273" type="SIM" status="Available" />
        </tbody>
      </table>
    </div>
  );
}

function InventoryCard({ title, count, status }: any) {
  return (
    <div className="p-6 bg-white border rounded-lg shadow-sm">
      <h3 className="text-slate-500 font-medium">{title}</h3>
      <p className="text-3xl font-bold mt-2">{count}</p>
    </div>
  );
}

function ResourceRow({ id, name, type, status }: any) {
  return (
    <tr className="border-t hover:bg-slate-50 transition">
      <td className="p-4 font-mono text-sm">{id}</td>
      <td className="p-4 font-medium">{name}</td>
      <td className="p-4">{type}</td>
      <td className="p-4">
        <span className="px-2 py-1 bg-green-100 text-green-700 rounded-full text-xs font-bold">{status}</span>
      </td>
      <td className="p-4">
        <button className="text-blue-600 hover:underline">Manage</button>
      </td>
    </tr>
  );
}
