import { useState } from 'react';
import { Section, Button} from '@telegram-apps/telegram-ui';

interface Paginator {
    limit: number;
    offset: number;
    total: number;
    page: number;
}
  
  export const Pagination: React.FC<{ paginator: Paginator }> = ({ paginator }) => {
  const [currentPage, setCurrentPage] = useState<number>(paginator.page);
  const pages = Math.ceil(paginator.total / paginator.limit);

  const handlePageChange = (newPage: number) => {
    setCurrentPage(newPage);
  };

  return (
    <Section>
      <div>
        <Button
          onClick={() => handlePageChange(currentPage - 1)}
          disabled={currentPage === 1}
        >
          Назад
        </Button>
        <Button 
            style={{ width: '100px' }} 
            onClick={() => handlePageChange(currentPage + 1)}
            disabled={currentPage == pages}
        >
          {currentPage} из {pages}
        </Button>
        <Button
          onClick={() => handlePageChange(currentPage + 1)}
          disabled={currentPage >= pages}
        >
          Дальше
        </Button>
      </div>
    </Section>
  );
};