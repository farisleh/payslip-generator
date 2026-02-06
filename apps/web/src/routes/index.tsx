import { createFileRoute } from "@tanstack/react-router";
import { useState } from "react";
import { getApiUrl } from "../config";

export const Route = createFileRoute("/")({
  component: Index,
});

interface PayslipResponse {
  employee_name: string;
  gross_monthly_income: string;
  monthly_income_tax: string;
  net_monthly_income: string;
}

function Index() {
  const [employeeName, setEmployeeName] = useState("");
  const [annualSalary, setAnnualSalary] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");
  const [result, setResult] = useState<PayslipResponse | null>(null);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    setError("");
    setResult(null);

    try {
      const response = await fetch(getApiUrl("/payslip"), {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          employee_name: employeeName,
          annual_salary: annualSalary,
        }),
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || "Failed to generate payslip");
      }

      const data = await response.json();
      setResult(data);
      setEmployeeName("");
      setAnnualSalary("");
    } catch (err) {
      setError(err instanceof Error ? err.message : "An error occurred");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
      <div className="px-4 py-6 sm:px-0">
        <div className="max-w-md mx-auto">
          <div className="bg-white shadow rounded-lg p-6">
            <h2 className="text-2xl font-bold mb-6 text-gray-900">
              Generate Payslip
            </h2>

            <form onSubmit={handleSubmit} className="space-y-4">
              <div>
                <label
                  htmlFor="employee-name"
                  className="block text-sm font-medium text-gray-700"
                >
                  Employee Name
                </label>
                <input
                  type="text"
                  id="employee-name"
                  value={employeeName}
                  onChange={(e) => setEmployeeName(e.target.value)}
                  required
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm px-3 py-2 border"
                  placeholder="Enter employee name"
                />
              </div>

              <div>
                <label
                  htmlFor="annual-salary"
                  className="block text-sm font-medium text-gray-700"
                >
                  Annual Salary
                </label>
                <input
                  type="number"
                  id="annual-salary"
                  value={annualSalary}
                  onChange={(e) => setAnnualSalary(e.target.value)}
                  required
                  min="0"
                  step="0.01"
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm px-3 py-2 border"
                  placeholder="Enter annual salary"
                />
              </div>

              {error && (
                <div className="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded">
                  {error}
                </div>
              )}

              <button
                type="submit"
                disabled={loading}
                className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 disabled:opacity-50 disabled:cursor-not-allowed"
              >
                {loading ? "Generating..." : "Generate Payslip"}
              </button>
            </form>

            {result && (
              <div className="mt-6 border-t pt-6">
                <h3 className="text-lg font-semibold mb-4 text-gray-900">
                  Payslip Details
                </h3>
                <dl className="space-y-2">
                  <div className="flex justify-between">
                    <dt className="text-sm text-gray-600">Employee:</dt>
                    <dd className="text-sm font-medium text-gray-900">
                      {result.employee_name}
                    </dd>
                  </div>
                  <div className="flex justify-between">
                    <dt className="text-sm text-gray-600">
                      Gross Monthly Income:
                    </dt>
                    <dd className="text-sm font-medium text-gray-900">
                      ${parseFloat(result.gross_monthly_income).toFixed(2)}
                    </dd>
                  </div>
                  <div className="flex justify-between">
                    <dt className="text-sm text-gray-600">Monthly Tax:</dt>
                    <dd className="text-sm font-medium text-gray-900">
                      ${parseFloat(result.monthly_income_tax).toFixed(2)}
                    </dd>
                  </div>
                  <div className="flex justify-between pt-2 border-t">
                    <dt className="text-sm font-semibold text-gray-900">
                      Net Monthly Income:
                    </dt>
                    <dd className="text-sm font-bold text-indigo-600">
                      ${parseFloat(result.net_monthly_income).toFixed(2)}
                    </dd>
                  </div>
                </dl>
              </div>
            )}
          </div>

          <div className="mt-6 bg-blue-50 border border-blue-200 rounded-lg p-4">
            <h3 className="text-sm font-semibold text-blue-900 mb-2">
              Tax Brackets
            </h3>
            <table className="min-w-full text-xs text-blue-900">
              <thead>
                <tr>
                  <th className="text-left py-1">Salary Range</th>
                  <th className="text-right py-1">Tax Rate</th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  <td className="py-1">$0 - $20,000</td>
                  <td className="text-right">0%</td>
                </tr>
                <tr>
                  <td className="py-1">$20,001 - $40,000</td>
                  <td className="text-right">10%</td>
                </tr>
                <tr>
                  <td className="py-1">$40,001 - $80,000</td>
                  <td className="text-right">20%</td>
                </tr>
                <tr>
                  <td className="py-1">$80,001 - $180,000</td>
                  <td className="text-right">30%</td>
                </tr>
                <tr>
                  <td className="py-1">$180,001+</td>
                  <td className="text-right">40%</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  );
}
